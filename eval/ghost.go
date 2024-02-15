package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

var Ghost types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return "ghost", nil
}
