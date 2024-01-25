package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func Lambda(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		if len(arguments) != len(trajectory.Expression.Parameters) {
			return nil, errors.New("Could not evaluate lambda, wrong number of arguments, expected " + fmt.Sprint(len(trajectory.Expression.Parameters)) + " got " + fmt.Sprint(len(arguments)))
		}

		scope := types.NewTrajectory(trajectory, trajectory.Expression)

		for i, param := range trajectory.Expression.Parameters {
			types.DefineName(&scope, param, arguments[i])
		}

		var value types.Value

		for _, exp := range trajectory.Expression.Children {
			child := types.NewTrajectory(&scope, exp)

			val, err := eval(&child)

			if err != nil {
				return nil, err
			}

			value = val
		}

		return value, nil
	}

	return lambda, nil
}

func Call(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	fnVal, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	fn, ok := fnVal.(types.Lambda)

	if !ok {
		return nil, errors.New("Error evaluating call, expression that didn't evaluate to a Lambda, got " + fmt.Sprint(fnVal))
	}

	children := trajectory.Children[1:]

	args := []types.Value{}

	if len(children) > 0 {
		for _, traj := range children {
			arg, err := eval(traj)

			if err != nil {
				return nil, err
			}

			args = append(args, arg)
		}
	}

	val, err := fn(args...)

	if err != nil {
		return nil, err
	}

	return val, nil
}
