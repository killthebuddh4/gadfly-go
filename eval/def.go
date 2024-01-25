package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
	traj "github.com/killthebuddh4/gadflai/types"
)

func Def(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	identifier := trajectory.Children[0].Expression.Operator.Value

	var value types.Value

	for _, input := range trajectory.Children[1:] {
		val, err := eval(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	traj.DefineName(trajectory.Parent, identifier, value)

	return value, nil
}

func Let(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	identifier := trajectory.Children[0].Expression.Operator.Value

	var value types.Value

	for _, input := range trajectory.Children[1:] {
		val, err := eval(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	traj.EditName(trajectory.Parent, identifier, value)

	return value, nil
}

func Identifier(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	if trajectory.Parent == nil {
		return nil, errors.New("cannot evaluate identifier " + trajectory.Expression.Operator.Value + " with nil parent")
	}
	return traj.ResolveName(trajectory.Parent, trajectory.Expression.Operator.Value)
}
