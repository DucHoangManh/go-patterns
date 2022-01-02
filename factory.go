package gopatterns

import (
	"errors"
	"fmt"
)

// use Stringer as the interface

type ErrPrint struct{}

func (p *ErrPrint) String() string {
	return "some error happens"
}

type InfoPrint struct{}

func (p *InfoPrint) String() string {
	return "nothing dramatically happens"
}

func NewPrinter(kind string) (result fmt.Stringer, err error) {
	switch kind {
	case "error":
		result = &ErrPrint{}
	case "info":
		result = &InfoPrint{}
	default:
		err = errors.New("invalid kind")
	}
	return
}
