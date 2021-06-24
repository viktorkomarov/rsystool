package main

type ActionType int

const (
	RegularFile ActionType = iota
	NamedPipes
	Terminal
	RemoteMachine
	ListOfUser
)

type Action struct {
	Type    ActionType
	Payload string
}

func ActionParse(action string) Action {
	return Action{}
}
