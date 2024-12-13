package m1errors

import (
	"fmt"
)

type Error struct {
	source
	Description error
}

func (e Error) Error() string {
	return fmt.Sprintf("source: %s, description: %d", e.source, e.Description)
}

func ParseReturnCode(code uint32) error {
	if code == ErrorOK {
		return nil
	}

	return Error{
		source:      source(code),
		Description: description(code),
	}
}
