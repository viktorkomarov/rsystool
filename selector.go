package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrIllegalSelector error = errors.New("illegal selector")
)

type Selector struct {
	Facilities map[Facility][]Priority
	ActionType ActionType
	ActionItem string
}

// Selector: Facility[,Facility].([[!]=]|[!])Priority
// Selector[;]Selector...
func (s *Selector) Parse(line string) error {
	if s.Facilities == nil {
		s.Facilities = make(map[Facility][]Priority)
	}

	for sep := strings.Index(line, ";"); sep != -1; {
		current := line[:sep]

		dot := strings.Index(current, ".")
		if dot == -1 {
			return fmt.Errorf("%w: %s", ErrIllegalSelector, line)
		}

		facilities, err := FacilityParse(current[:dot])
		if err != nil {
			return err
		}

		policy, err := PriorityParse(current[dot+1:])
		if err != nil {
			return err
		}

		s.applyPolicy(facilities, policy)
		line = line[sep+1:]
	}

	return nil
}

// can be !* or =* think about it
func (s *Selector) applyPolicy(facilities []Facility, policy PriorityPolicy) {
	if policy.Priority == None {
		for _, fac := range facilities {
			delete(s.Facilities, fac)
		}
	}

	if policy.OnlyEqual && !policy.Negative {
		for _, fac := range facilities {
			s.Facilities[fac] = []Priority{policy.Priority}
		}

		return
	}

	if policy.Negative && policy.OnlyEqual {
		for _, fac := range facilities {
			s.Facilities[fac] = excludePriority(s.Facilities[fac], policy.Priority)
		}

		return
	}

	for _, fac := range facilities {
		s.Facilities[fac] = collectPriorities(s.Facilities[fac], policy.Negative, policy.Priority)
	}
}

func collectPriorities(existing []Priority, negative bool, priority Priority) []Priority {
	if priority == AllPriority {
		return prioritiesCopy()
	}

	if negative {
		for i, p := range existing {
			if p == priority {
				return existing[:i]
			}
		}

		return existing
	}

	priorities := prioritiesCopy()
	for i, prior := range priorities {
		if prior == priority {
			return priorities[:i+1]
		}
	}

	return priorities
}

func excludePriority(priorities []Priority, p Priority) []Priority {
	i := -1
	for ; i < len(priorities); i++ {
		if priorities[i] == p {
			break
		}
	}

	if i == -1 {
		return priorities
	}

	return append(priorities[:i], priorities[i+1:]...)
}
