package error

import (
	"fmt"

	"github.com/small-entropy/go-backbone/stores/abstract"
)

type HandlerError struct {
	Status  string
	Code    int
	Message string
	Err     error
}

func (he *HandlerError) Error() string {
	return fmt.Sprintf("Status code: %d.\nStatus: %v.Message: %v.\nError: %v", he.Code, he.Status, he.Message, he.Err)
}

type ControllerError[DATA any] struct {
	Status  string
	Message string
	Data    *DATA
	Filter  *map[string]interface{}
	Update  *map[string]interface{}
	Page    *abstract.Page
	Err     error
}

func (ce *ControllerError[DATA]) Error() string {
	return fmt.Sprintf("Status: %v.\nMessage: %v.\nError: %v", ce.Status, ce.Message, ce.Err)
}

type StoreError struct {
	Status       string
	StorageName  string
	DatabaseName string
	Err          error
}

func (se *StoreError) Error() string {
	return fmt.Sprintf("Status: %v.\nStorage name: %v.\nError: %v", se.Status, se.StorageName, se.Err)
}
