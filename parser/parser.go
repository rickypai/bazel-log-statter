package parser

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rickypai/bazel-log-statter/bazel"
)

var (
	cachedLineRegex   = regexp.MustCompile(`(?P<target>\/\/.+)\s+(?P<cached>\(cached\))\s+(?P<status>PASSED)\s+in\s+(?P<duration>.+)s`)
	uncachedLineRegex = regexp.MustCompile(`(?P<target>\/\/.+)\s+(?P<status>PASSED|FAILED)\s+in\s+(?P<duration>.+)s`)
	noStatusLineRegex = regexp.MustCompile(`(?P<target>\/\/.+)\s+(?P<status>NO\sSTATUS)`)
	flakyLineRegex    = regexp.MustCompile(`(?P<target>\/\/.+)\s+(?P<status>FLAKY),\sfailed\sin\s(?P<success>\d+)\sout\sof\s(?P<tries>\d+)\sin\s+(?P<duration>.+)s`)
)

func ParseLine(line string) (result *bazel.TargetResult, err error) {
	result, err = cachedMatches(line)
	if result != nil {
		return
	}

	result, err = uncachedMatches(line)
	if result != nil {
		return
	}

	result, err = noStatusMatches(line)
	if result != nil {
		return
	}

	result, err = flakyMatches(line)
	if result != nil {
		return
	}

	return nil, nil
}

func cachedMatches(line string) (*bazel.TargetResult, error) {
	var err error

	matches := cachedLineRegex.FindStringSubmatch(line)

	if len(matches) == 0 {
		return nil, nil
	}

	result := &bazel.TargetResult{}
	result.Name = strings.TrimSpace(matches[1])
	result.Cached = matches[2] == "(cached)"
	result.Status = bazel.Status(matches[3])
	result.Duration, err = parseDuration(matches[4])
	if err != nil {
		return nil, err
	}

	return result, nil
}

func flakyMatches(line string) (*bazel.TargetResult, error) {
	var err error

	matches := flakyLineRegex.FindStringSubmatch(line)

	if len(matches) == 0 {
		return nil, nil
	}

	result := &bazel.TargetResult{}
	result.Name = strings.TrimSpace(matches[1])
	result.Status = bazel.Status(matches[2])
	result.Successes, err = strconv.Atoi(matches[3])
	if err != nil {
		return nil, err
	}
	result.Attempts, err = strconv.Atoi(matches[4])
	if err != nil {
		return nil, err
	}
	result.Duration, err = parseDuration(matches[5])
	if err != nil {
		return nil, err
	}

	return result, nil
}

func uncachedMatches(line string) (*bazel.TargetResult, error) {
	var err error

	matches := uncachedLineRegex.FindStringSubmatch(line)

	if len(matches) == 0 {
		return nil, nil
	}

	result := &bazel.TargetResult{}
	result.Name = strings.TrimSpace(matches[1])
	result.Status = bazel.Status(matches[2])
	result.Duration, err = parseDuration(matches[3])
	if err != nil {
		return nil, err
	}

	return result, nil
}

func noStatusMatches(line string) (*bazel.TargetResult, error) {
	matches := noStatusLineRegex.FindStringSubmatch(line)

	if len(matches) == 0 {
		return nil, nil
	}

	result := &bazel.TargetResult{}
	result.Name = strings.TrimSpace(matches[1])
	result.Status = bazel.Status(matches[2])

	return result, nil
}

func parseDuration(durationStr string) (time.Duration, error) {
	durationF, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(durationF * float64(time.Second)), nil
}
