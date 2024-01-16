package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Def(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	identifier := trajectory.Children[0].Expression.Operator.Value

	var value value.Value

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

func Let(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	identifier := trajectory.Children[0].Expression.Operator.Value

	var value value.Value

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

func Identifier(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	if trajectory.Parent == nil {
		return nil, errors.New("cannot evaluate identifier " + trajectory.Expression.Operator.Value + " with nil parent")
	}
	return traj.ResolveName(trajectory.Parent, trajectory.Expression.Operator.Value)
}
