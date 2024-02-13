package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

func SchemaIdentity() types.Lambda {
	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		return arguments[0], nil
	}
	return lambda
}
