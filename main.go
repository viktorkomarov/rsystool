package main

import "unicode"

type RsyslogFormat int

const (
	Basic RsyslogFormat = iota
	Advanced
	ObsoleteLegacy
)

type Selector struct {
	Facility string // insensetive
	Priority string // insensetive
}

func isContainOnlyDigit(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func main() {
	// write test TDD
	// to int for easire ordering
}
