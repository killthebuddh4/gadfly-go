package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

func SchemaIdentity() types.Exec {
	var lambda types.Exec = func(arguments ...types.Value) (types.Value, error) {
		return arguments[0], nil
	}
	return lambda
}
