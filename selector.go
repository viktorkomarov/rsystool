package main

type Selector struct {
	Facilities []Facility
	Priorities []Priority
	ActionType ActionType
	ActionItem string
}

// Selector: Facility[,Facility].([[!]=]|[!])Priority
// Selector[;]Selector...
func ParseSelector(line string) (Selector, error) {
	return Selector{}, nil
}
