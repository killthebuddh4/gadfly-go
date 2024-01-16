package eval

import (
	"errors"
	"fmt"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Lambda(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	var lambda traj.Lambda = func(arguments ...value.Value) (value.Value, error) {
		if len(arguments) != len(trajectory.Expression.Parameters) {
			return nil, errors.New("Could not evaluate lambda, wrong number of arguments, expected " + fmt.Sprint(len(trajectory.Expression.Parameters)) + " got " + fmt.Sprint(len(arguments)))
		}

		scope := traj.Traj(trajectory, trajectory.Expression)

		for i, param := range trajectory.Expression.Parameters {
			traj.DefineName(&scope, param, arguments[i])
		}

		var value value.Value

		for _, exp := range trajectory.Expression.Children {
			child := traj.Traj(&scope, exp)

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

func Call(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	fnVal, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	fn, ok := fnVal.(traj.Lambda)

	if !ok {
		return nil, errors.New("Error evaluating call, expression that didn't evaluate to a Lambda, got " + fmt.Sprint(fnVal))
	}

	children := trajectory.Children[1:]

	args := []value.Value{}

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
