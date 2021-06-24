package main

import "strconv"

type Priority int

const (
	Emerg Priority = iota
	Alert
	Crit
	Err
	Warning
	Notice
	Info
	Debug
	AllPriority
	NonePriority
)

var priorityOneOf = map[Priority]bool{
	Emerg: true, Alert: true, Crit: true,
	Err: true, Warning: true, Notice: true,
	Info: true, Debug: true,
}

var priorityName = map[string]Priority{
	"emerg":   Emerg,
	"alert":   Alert,
	"crit":    Crit,
	"debug":   Debug,
	"info":    Info,
	"notice":  Notice,
	"warn":    Warning,
	"warning": Warning,
	"err":     Err,
	"*":       AllPriority,
}

func PriorityParse(priority string) Priority {
	// check ([!]=)|(!)
	if isContainOnlyDigit(priority) {
		num, err := strconv.Atoi(priority)
		if err != nil {
			return NonePriority
		}

		prior := Priority(num)
		if priorityOneOf[prior] {
			return prior
		}

		return NonePriority
	}

	if p, ok := priorityName[priority]; ok {
		return p
	}

	return NonePriority
}
