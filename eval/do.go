package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

var Do types.Exec = func(arguments ...types.Value) (types.Value, error) {
	return arguments[len(arguments)-1], nil
}
