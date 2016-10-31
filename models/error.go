package models

import (
	"fmt"
)

//
type ErrNavNotExist struct {
	Errstr string
}

//
func IsErrNavNotExist(err error) bool {
	_, ok := err.(ErrNavNotExist)
	return ok
}

func (err ErrNavNotExist) Error() string {
	return fmt.Sprintf(err.Errstr)
}
