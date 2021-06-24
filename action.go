package main

type ActionType int

const (
	RegularFile ActionType = iota
	NamedPipes
	Terminal
	RemoteMachine
	ListOfUser
	UnknownAction
)

func ActionParse(action string) (ActionType, string) {
	return UnknownAction, ""
}
