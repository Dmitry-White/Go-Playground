package dal

import "fmt"

type User struct {
	Login   string
	Age     int
	Address string
}

type CustomError struct {
	Reason string
	Info   interface{}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s (%#v)", e.Reason, e.Info)
}
