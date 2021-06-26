package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrIllegalPriority error = errors.New("illegal priority")
)

type PriorityPolicy struct {
	Negative  bool
	OnlyEqual bool
	Priority  Priority
}

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
	None
	AllPriority
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
	"none":    None,
	"*":       AllPriority,
}

func prioritiesCopy() []Priority {
	return []Priority{Emerg, Alert, Crit, Err, Warning, Notice, Info, Debug}
}

func PriorityParse(line string) (PriorityPolicy, error) {
	negative := false
	if line[0] == '!' {
		line = line[1:]
		negative = true
	}

	onlyEqual := false
	if line[0] == '=' {
		line = line[1:]
		onlyEqual = true
	}

	var priority Priority
	if isContainOnlyDigit(line) {
		num, err := strconv.Atoi(line)
		if err != nil {
			return PriorityPolicy{}, fmt.Errorf("%w: %s", ErrIllegalPriority, err)
		}

		p := Priority(num)
		if !priorityOneOf[p] {
			return PriorityPolicy{}, fmt.Errorf("%w: unknown %d", ErrIllegalPriority, num)
		}

		priority = p
	} else {
		if p, ok := priorityName[line]; !ok {
			return PriorityPolicy{}, fmt.Errorf("%w: unknown %s", ErrIllegalPriority, line)
		} else {
			priority = p
		}
	}

	return PriorityPolicy{
		Negative:  negative,
		OnlyEqual: onlyEqual,
		Priority:  priority,
	}, nil
}
