package main

type Command struct {
	Command string
	F       func([]string) error
	Man     string
}

type WrongArgumentsError struct {
	msg string
}

func (e *WrongArgumentsError) Error() string { return e.msg }
