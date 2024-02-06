package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func If(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	err := types.ExpandBy(trajectory, trajectory.Expression.Children[0])

	if err != nil {
		return nil, err
	}

	whenExp := trajectory.Children[0]

	conditionVal, err := eval(whenExp)

	if err != nil {
		return nil, err
	}

	condition, ok := conditionVal.(bool)

	if !ok {
		return nil, errors.New("condition is not a boolean")
	}

	if condition {
		err := types.ExpandBy(trajectory, trajectory.Expression.Children[1])
		if err != nil {
			return nil, err
		}
	} else {
		err := types.ExpandBy(trajectory, trajectory.Expression.Children[2])
		if err != nil {
			return nil, err
		}
	}
	exp := trajectory.Children[1]
	return eval(exp)
}

func And(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	if (len(trajectory.Expression.Children) % 2) != 0 {
		return nil, errors.New("and must have even number of inputs")
	}

	var val types.Value = nil

	for i := 0; i < len(trajectory.Expression.Children); i += 2 {
		types.ExpandBy(trajectory, trajectory.Expression.Children[i])
		conditionVal, err := eval(trajectory.Children[i])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if !condition {
			return false, nil
		}

		types.ExpandBy(trajectory, trajectory.Expression.Children[i+1])
		body, err := eval(trajectory.Children[i+1])

		if err != nil {
			return nil, err
		}

		val = body
	}

	return val, nil
}

func Or(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	if (len(trajectory.Expression.Children) % 2) != 0 {
		return nil, errors.New("or must have even number of inputs")
	}

	for i := 0; i < len(trajectory.Expression.Children); i += 2 {
		types.ExpandBy(trajectory, trajectory.Expression.Children[i])
		conditionVal, err := eval(trajectory.Children[len(trajectory.Children)-1])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if condition {
			types.ExpandBy(trajectory, trajectory.Expression.Children[i+1])
			return eval(trajectory.Children[len(trajectory.Children)-1])
		}
	}

	return nil, nil
}

func While(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandBy(trajectory, trajectory.Expression.Children[0])

	condV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	cond, ok := condV.(bool)

	if !ok {
		return nil, errors.New("not a boolean")
	}

	if !cond {
		return nil, nil
	}

	var value types.Value = nil

	types.ExpandBy(trajectory, trajectory.Expression.Children[1])
	fnV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(types.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	for {
		condV, err := eval(trajectory.Children[0])

		if err != nil {
			return nil, err
		}

		cond, ok := condV.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if !cond {
			break
		}

		val, err := fn()

		if err != nil {
			return nil, err
		}

		value = val
	}

	return value, nil
}
