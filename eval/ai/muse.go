package ai

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func Muse(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	var gadfly types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		if len(arguments) != len(trajectory.Expression.Parameters) {
			return nil, errors.New("Could not evaluate lambda, wrong number of arguments, expected " + fmt.Sprint(len(trajectory.Expression.Parameters)) + " got " + fmt.Sprint(len(arguments)))
		}

		scope := types.NewTrajectory(trajectory, nil)

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

	return gadfly, nil
}
