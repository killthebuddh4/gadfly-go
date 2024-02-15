package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func Lambda(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	var lambda types.Exec = func(arguments ...types.Value) (types.Value, error) {
		if len(arguments) != len(trajectory.Expression.Parameters) {
			fmt.Println(arguments)
			return nil, errors.New("Could not evaluate lambda, wrong number of arguments, expected " + fmt.Sprint(len(trajectory.Expression.Parameters)) + " got " + fmt.Sprint(len(arguments)))
		}

		scope := types.NewTrajectory(trajectory, nil)

		for i, param := range trajectory.Expression.Parameters {
			// TODO ! HACK HACK HACK OH NO MY EYES
			types.DefineName(&scope, param.Children[0].Operator.Value, arguments[i])
		}

		for _, child := range trajectory.Expression.Parameters {
			validationTrajectory := types.NewTrajectory(&scope, child)

			eval, dispatchErr := dispatch(&validationTrajectory)

			if dispatchErr != nil {
				return nil, dispatchErr
			}

			_, err := eval(&validationTrajectory, Exec)

			if err != nil {
				return nil, err
			}
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

		if len(trajectory.Expression.Returns) == 0 {
			return value, nil
		} else {
			validationTrajectory := types.NewTrajectory(trajectory, trajectory.Expression.Returns[0])
			eval, dispatchErr := dispatch(&validationTrajectory)

			if dispatchErr != nil {
				return nil, dispatchErr
			}

			schemaV, err := eval(&validationTrajectory, Exec)

			if err != nil {
				return nil, err
			}

			schema, ok := schemaV.(types.Exec)

			if !ok {
				return nil, fmt.Errorf("not a function")
			}

			return schema(value)
		}
	}

	return lambda, nil
}

func Call(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	fnVal, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	fn, ok := fnVal.(types.Exec)

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
