package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

func Symbol(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	return trajectory.Expression.Operator.Value, nil
}
