package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

var Gadfly types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return "gadfly", nil
}
