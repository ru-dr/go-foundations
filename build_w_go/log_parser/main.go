/*
  2026-08-24T10:15:02 INFO server listening on :8080
  2026-08-24T10:15:09 WARN disk usage at 81%
  2026-08-24T10:16:44 ERROR failed to connect to postgres
  2026-08-24T10:16:45 INFO retrying connection
  garbage line with no level
*/

/*
 * OUTPUT

 INFO: 2
 WARN: 1
 ERROR: 1
 skipped 1 malformed line
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type logLine struct {
	date  string
	level string
	msg   string
}

func parseLine(line string) (logLine, error) {
	parts := strings.Fields(line)

	if len(parts) < 3 {
		return logLine{}, errors.New("garbage value found")
	}

	date := parts[0]
	level := parts[1]
	msg := strings.Join(parts[2:], " ")

	return logLine{date: date, level: level, msg: msg}, nil
}

func main() {
	file, err := os.Open("log.txt")
	defer file.Close()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(nil, err)
		defer file.Close()
		return
	}

	logCounter := make(map[string]int)
	skipped := 0

	for _, line := range lines {
		parsed, err := parseLine(line)

		if err != nil {
			fmt.Println(err)
			skipped++
			continue
		}
		_, ok := logCounter[parsed.level]
		if !ok {
			logCounter[parsed.level] = 1
		} else {
			logCounter[parsed.level]++
		}
		fmt.Println(parsed)

	}
	for k, v := range logCounter {
		fmt.Println(k, v)
	}
	fmt.Println("skipped", skipped, "malformed lines")
}
