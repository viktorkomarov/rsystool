package main

type Selector struct {
	Facilities []Facility
	Priorities []Priority
}

func ParseSelector(line string) (Selector, error) {
	return Selector{}, nil
}
