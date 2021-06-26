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

	for sep := strings.Index(line, ";"); sep != -1; sep = strings.Index(line, ";") {
		if err := s.parse(line[:sep]); err != nil {
			return err
		}

		line = line[sep+1:]
	}

	return s.parse(line)
}

func (s *Selector) parse(current string) error {
	dot := strings.Index(current, ".")
	if dot == -1 {
		return fmt.Errorf("%w: %s", ErrIllegalSelector, current)
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
	return nil
}

// can be !* or =* think about it
func (s *Selector) applyPolicy(facilities []Facility, policy PriorityPolicy) {
	fmt.Printf("%+v\n", s.Facilities)
	if policy.Priority == None {
		for _, fac := range facilities {
			delete(s.Facilities, fac)
		}

		return
	}

	if policy.OnlyEqual && !policy.Negative {
		for _, fac := range facilities {
			s.Facilities[fac] = append(s.Facilities[fac], policy.Priority)
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
		s.Facilities[fac] = collectPriorities(s.Facilities[fac], policy)
	}
}

func collectPriorities(existing []Priority, policy PriorityPolicy) []Priority {
	if policy.Priority == AllPriority {
		return prioritiesCopy()
	}

	if policy.Negative {
		for i := len(existing) - 1; i >= 0; i-- {
			if existing[i] == policy.Priority {
				return existing[i+1:]
			}
		}

		return existing
	}

	priorities := prioritiesCopy()
	for i, prior := range priorities {
		if prior == policy.Priority {
			return priorities[:i+1]
		}
	}

	return priorities
}

func excludePriority(priorities []Priority, p Priority) []Priority {
	i := -1
	for j := 0; j < len(priorities); j++ {
		if priorities[j] == p {
			i = j
		}
	}

	if i == -1 {
		return priorities
	}

	return append(priorities[:i], priorities[i+1:]...)
}
