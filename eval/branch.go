package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func If(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	whenExp := trajectory.Children[0]
	thenExp := trajectory.Children[1]
	elseExp := trajectory.Children[2]

	conditionVal, err := eval(whenExp)

	if err != nil {
		return nil, err
	}

	condition, ok := conditionVal.(bool)

	if !ok {
		return nil, errors.New("condition is not a boolean")
	}

	if condition {
		return eval(thenExp)
	} else {
		return eval(elseExp)
	}
}

func And(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	if (len(trajectory.Children) % 2) != 0 {
		return nil, errors.New("and must have even number of inputs")
	}

	var val value.Value = nil

	for i := 0; i < len(trajectory.Children); i += 2 {
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

		body, err := eval(trajectory.Children[i+1])

		if err != nil {
			return nil, err
		}

		val = body
	}

	return val, nil
}

func Or(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	if (len(trajectory.Children) % 2) != 0 {
		return nil, errors.New("or must have even number of inputs")
	}

	for i := 0; i < len(trajectory.Children); i += 2 {
		conditionVal, err := eval(trajectory.Children[i])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if condition {
			return eval(trajectory.Children[i+1])
		}
	}

	return nil, nil
}

func While(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	var value value.Value = nil

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

		for _, child := range trajectory.Children[1:] {
			val, err := eval(child)

			if err != nil {
				return nil, err
			}

			value = val
		}
	}

	return value, nil
}
