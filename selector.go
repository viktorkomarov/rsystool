package main

import "github.com/looplab/fsm"

type Selector struct {
	Facilities []Facility
	Priorities []Priority
	ActionType ActionType
	ActionItem string
}

// SelectorParts
const (
	from     = "init"
	facility = "facility"
	priority = "priority"
	terminal = "terminal"
)

// Selector: Facility[,Facility].([[!]=]|[!])Priority
// Selector[;]Selector...
func ParseSelector(line string) (Selector, error) {
	stateMachine := fsm.NewFSM(from, []fsm.EventDesc{
		{Name: from, Src: []string{from}, Dst: facility},
		{Name: ".", Src: []string{facility}, Dst: priority},
		{Name: ";", Src: []string{facility}, Dst: priority},
		{Name: "", Src: []string{priority}, Dst: terminal},
	}, nil)

	return Selector{}, nil
}
