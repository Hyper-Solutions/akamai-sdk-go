package akamai

import (
	"errors"
	"io"
	"regexp"
	"strconv"
)

var (
	secDurationExpr = regexp.MustCompile(`data-duration=(\d+)`)
	secPageExpr     = regexp.MustCompile(`src="(/_sec/cp_challenge/ak-challenge-\d+-\d+.htm)"`)

	ErrSecCpt = errors.New("akamai-sdk-go: error parsing sec-cpt page")
)

// ParseSecCpt parses the duration of a sec-cpt challenge and the path of the page that contains the challenge itself
func ParseSecCpt(reader io.Reader) (int, string, error) {
	src, err := io.ReadAll(reader)
	if err != nil {
		return 0, "", errors.Join(ErrSecCpt, err)
	}

	durationMatches := secDurationExpr.FindSubmatch(src)
	if len(durationMatches) < 2 {
		return 0, "", ErrPixelHtmlVarNotFound
	}

	duration, err := strconv.Atoi(string(durationMatches[1]))
	if err != nil {
		return 0, "", errors.Join(ErrSecCpt, err)
	}

	pageMatches := secPageExpr.FindSubmatch(src)
	if len(pageMatches) < 2 {
		return 0, "", ErrPixelHtmlVarNotFound
	}

	return duration, string(pageMatches[1]), nil
}
