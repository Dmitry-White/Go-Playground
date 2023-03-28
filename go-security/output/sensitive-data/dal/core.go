package dal

import (
	"fmt"
	"log"
)

type User struct {
	Login   string
	Age     int
	Address string
}

type CustomError struct {
	Reason string
	Info   interface{}
}

type AuditLog struct {
	Origin string
	Data   interface{}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s (%#v)", e.Reason, e.Info)
}

func (a *AuditLog) Track() {
	// Dummy audit log
	log.Printf("Audit: \n%v from %s\n", a.Data, a.Origin)
}
