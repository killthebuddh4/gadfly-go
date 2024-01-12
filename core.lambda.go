package main

type Lambda func(scope *Trajectory) (Value, error)

func EvaluateLambda(closure *Trajectory) (Value, error) {
	var lambda Lambda = func(scope *Trajectory) (Value, error) {
		//
		// EVALUATE ARGUMENTS
		//

		expand(scope)

		namespace := Traj(closure, closure.Expression)

		for i, child := range scope.Children {
			arg, err := Evaluate(child)

			if err != nil {
				return nil, err
			}

			param := closure.Expression.Parameters[i]

			DefineName(closure, param, func(_ *Trajectory) (Value, error) {
				return arg, nil
			})
		}

		//
		// EVALUATE FUNCTION BODY
		//

		var value Value

		for _, child := range closure.Expression.Children {
			traj := Traj(&namespace, child)

			val, err := Evaluate(&traj)

			if err != nil {
				return nil, err
			}

			value = val
		}

		return value, nil
	}

	return lambda, nil
}
