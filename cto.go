package cto

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	noColor = "\033[0m"
	red     = "\033[0;31m"
	green   = "\033[0;32m"
	yellow  = "\033[0;33m"
	blue    = "\033[0;34m"
	purple  = "\033[0;35m"
	cyan    = "\033[0;36m"
)

const (
	pass  = "--- PASS:"
	fail  = "--- FAIL:"
	debug = "DEBUG:"
)

// Colorize ..
func Colorize(r io.Reader, w io.Writer) error {
	var failed []string

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if isPass(line) {
			fmt.Fprintf(w, "%s%s%s\n", green, line, noColor)
			continue
		}

		if isFail(line) {
			fmt.Fprintf(w, "%s%s%s\n", red, line, noColor)
			failed = append(failed, line)
			continue
		}

		if isDebug(line) {
			fmt.Fprintf(w, "%s%s%s\n", blue, line, noColor)
			continue
		}

		if isUpdate(line) {
			fmt.Fprintf(w, "%s\n", line)
			continue
		}

		fmt.Fprintf(w, "%s%s%s\n", yellow, line, noColor)
	}

	if len(failed) > 0 {
		fmt.Fprintf(w, "\nFailed tests: %d\n", len(failed))
		for _, f := range failed {
			fmt.Fprintf(w,
				"%v\n\t%sgo test -race -v ./... -run \"%v\"%s\n",
				strings.TrimSpace(f), cyan, getTestName(f), noColor,
			)
		}
		return fmt.Errorf("%d tests failed", len(failed))
	}

	return nil
}

func isUpdate(line string) bool {
	return startsWith(line, "=== RUN") ||
		startsWith(line, "=== CONT") ||
		startsWith(line, "=== PAUSE") ||
		startsWith(line, "FAIL") ||
		startsWith(line, "?") ||
		startsWith(line, "PASS") ||
		startsWith(line, "ok ")
}

func isDebug(line string) bool {
	return startsWith(line, debug)
}

func isPass(line string) bool {
	return startsWith(strings.TrimSpace(line), pass)
}

func isFail(line string) bool {
	return startsWith(strings.TrimSpace(line), fail)
}

func startsWith(line, prefix string) bool {
	return len(line) >= len(prefix) && line[0:len(prefix)] == prefix
}

func getTestName(line string) string {
	parts := strings.Split(strings.TrimSpace(line), " ")
	return parts[2]
}
