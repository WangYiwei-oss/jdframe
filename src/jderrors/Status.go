package jderrors

type Status struct {
	Code Code
	Messaage string
}

func NewStatus(code Code, messaage string) *Status {
	return &Status{Code: code, Messaage: messaage}
}
