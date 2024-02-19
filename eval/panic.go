package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

var Panic types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	panic("Panic :: not implemented")
}
