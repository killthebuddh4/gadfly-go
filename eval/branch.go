package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func If(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
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

func And(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	cases := []CaseHandler{}

	for _, child := range trajectory.Children {
		caseHandlerV, err := eval(child)

		if err != nil {
			return nil, err
		}

		caseHandler, ok := caseHandlerV.(CaseHandler)

		if !ok {
			return nil, errors.New("not a case handler")
		}

		cases = append(cases, caseHandler)
	}

	var value types.Value = nil

	for _, caseHandler := range cases {
		condV, err := caseHandler.Cond()

		if err != nil {
			return nil, err
		}

		cond, ok := condV.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if !cond {
			return nil, nil
		} else {
			val, err := caseHandler.Body()

			if err != nil {
				return nil, err
			}

			value = val
		}
	}

	return value, nil
}

func Or(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	cases := []CaseHandler{}

	for _, child := range trajectory.Children {
		caseHandlerV, err := eval(child)

		if err != nil {
			return nil, err
		}

		caseHandler, ok := caseHandlerV.(CaseHandler)

		if !ok {
			return nil, errors.New("not a case handler")
		}

		cases = append(cases, caseHandler)
	}

	for _, caseHandler := range cases {
		condV, err := caseHandler.Cond()

		if err != nil {
			return nil, err
		}

		cond, ok := condV.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if cond {
			return caseHandler.Body()
		}
	}

	return nil, nil
}

func While(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
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
