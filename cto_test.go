package cto_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/slavsan/cto"
)

var escapeSequences = map[string]string{
	"\n":         "\\n",
	"\t":         "__TAB__",
	"\033[0m":    "__NO_COLOR__",
	"\033[0;31m": "__RED__",
	"\033[0;32m": "__GREEN__",
	"\033[0;33m": "__YELLOW__",
	"\033[0;34m": "__BLUE__",
	"\033[0;35m": "__PURPLE__",
	"\033[0;36m": "__CYAN__",
}

func escape(line string) string {
	for k, v := range escapeSequences {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}

func TestExamples(t *testing.T) {
	testCases := []struct {
		Title    string
		Example  string
		Expected string
	}{
		{
			Title: "With 1 passing test",
			Example: "" +
				"=== RUN   TestMyFunc\n" +
				"--- PASS: TestMyFunc (0.00s)\n" +
				"PASS\n" +
				"ok      github.com/slavsan/cto  0.243s\n" +
				"?       github.com/slavsan/cto/cmd      [no test files]\n",
			Expected: "" +
				"=== RUN   TestMyFunc\n" +
				"__GREEN__--- PASS: TestMyFunc (0.00s)__NO_COLOR__\n" +
				"PASS\n" +
				"ok      github.com/slavsan/cto  0.243s\n" +
				"?       github.com/slavsan/cto/cmd      [no test files]\n",
		},
		{
			Title: "With one failing test",
			Example: "" +
				"=== RUN   MyTestFunc\n" +
				"    cto_test.go:355:\n" +
				"                Error Trace:    cto_test.go:355\n" +
				"                Error:          Expected nil, but got: 3\n" +
				"                Test:           MyTestFunc\n" +
				"--- FAIL: MyTestFunc (0.00s)\n" +
				"FAIL\n" +
				"FAIL    github.com/slavsan/cto  0.464s\n" +
				"?       github.com/slavsan/cto/cmd      [no test files]\n" +
				"FAIL\n",
			Expected: "" +
				"=== RUN   MyTestFunc\n" +
				"__YELLOW__    cto_test.go:355:__NO_COLOR__\n" +
				"__YELLOW__                Error Trace:    cto_test.go:355__NO_COLOR__\n" +
				"__YELLOW__                Error:          Expected nil, but got: 3__NO_COLOR__\n" +
				"__YELLOW__                Test:           MyTestFunc__NO_COLOR__\n" +
				"__RED__--- FAIL: MyTestFunc (0.00s)__NO_COLOR__\n" +
				"FAIL\n" +
				"FAIL    github.com/slavsan/cto  0.464s\n" +
				"?       github.com/slavsan/cto/cmd      [no test files]\n" +
				"FAIL\n" +
				"\n" +
				"Failed tests: 1\n" +
				"--- FAIL: MyTestFunc (0.00s)\n" +
				"	__CYAN__go test -race -v ./... -run \"MyTestFunc\"__NO_COLOR__\n",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			tc.Title,
			func(t *testing.T) {
				var actual bytes.Buffer
				r := strings.NewReader(tc.Example)

				cto.Colorize(r, &actual)

				if escape(tc.Expected) != escape(actual.String()) {
					t.Errorf(
						"The two sides arent equal\n"+
							"\t\tExpected: '%s'\n"+
							"\t\t  Actual: '%s'\n",
						escape(tc.Expected),
						escape(actual.String()),
					)
				}
			})
	}
}
