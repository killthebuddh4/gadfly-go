package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

type Dispatch = func(*types.Expression) (types.Exec, error)

func Exec(scope *types.Trajectory, expr *types.Expression, d Dispatch) (types.Value, error) {
	trajectory := types.NewTrajectory(scope, expr)

	eval, err := d(expr)

	if err != nil {
		return nil, err
	}

	if expr.Operator.Type == "fn" {
		return eval, nil
	} else {
		args := []types.Value{}

		for _, child := range expr.Children {
			value, err := Exec(&trajectory, child, d)

			if err != nil {
				return nil, err
			}

			args = append(args, value)
		}

		return eval(args...)
	}
}

// func Exec(trajectory *types.Trajectory) (types.Value, error) {
// 	eval, dispatchErr := dispatch(trajectory)

// 	if dispatchErr != nil {
// 		return nil, dispatchErr
// 	}

// 	if trajectory.Expression.Operator.Value != "fn" {
// 		for _, child := range trajectory.Expression.Parameters {
// 			validationTrajectory := types.NewTrajectory(trajectory, child)

// 			eval, dispatchErr := dispatch(&validationTrajectory)

// 			if dispatchErr != nil {
// 				return nil, dispatchErr
// 			}

// 			_, err := eval(&validationTrajectory, Exec)

// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	val, evalErr := eval(trajectory, Exec)

// 	if evalErr != nil {
// 		return nil, evalErr
// 	}

// 	if trajectory.Expression.Operator.Value == "fn" {
// 		return val, nil
// 	} else {
// 		if len(trajectory.Expression.Returns) == 0 {
// 			return val, nil
// 		} else {
// 			validationTrajectory := types.NewTrajectory(trajectory, trajectory.Expression.Returns[0])
// 			eval, dispatchErr := dispatch(&validationTrajectory)

// 			if dispatchErr != nil {
// 				return nil, dispatchErr
// 			}

// 			schemaV, err := eval(&validationTrajectory, Exec)

// 			if err != nil {
// 				return nil, err
// 			}

// 			schema, ok := schemaV.(types.Lambda)

// 			if !ok {
// 				return nil, fmt.Errorf("not a function")
// 			}

// 			return schema(val)
// 		}
// 	}
// }
