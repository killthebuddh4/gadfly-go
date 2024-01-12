package main

import "fmt"

type Evaluator func(*Trajectory) (Value, error)

func Evaluate(trajectory *Trajectory) (Value, error) {
	if trajectory.Expression.Operator.Lexeme == "fn" {
		return EvaluateFn(trajectory)
	} else {
		switch trajectory.Expression.Operator.Type {
		case TOKENS.Root:
			return EvaluateRoot(trajectory)
		case TOKENS.String:
			return EvaluateString(trajectory)
		case TOKENS.Number:
			return EvaluateNumber(trajectory)
		default:
			return EvaluateName(trajectory)
		}
	}
}

func EvaluateRoot(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	for _, child := range trajectory.Children {
		_, err := Evaluate(child)

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func EvaluateName(trajectory *Trajectory) (Value, error) {
	fmt.Println("Evaluating name", trajectory.Expression.Operator.Lexeme)
	expand(trajectory)

	args := []Value{}

	for _, child := range trajectory.Children {
		val, err := Evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, val)
	}

	lambda, err := ResolveName(trajectory.Parent, trajectory.Expression.Operator.Lexeme)

	if err != nil {
		return nil, err
	}

	return lambda(trajectory.Parent, args...)
}

func EvaluateFn(trajectory *Trajectory) (Value, error) {
	fmt.Println("Evaluating fn")
	var lambda Lambda = func(scope *Trajectory, arguments ...Value) (Value, error) {
		fmt.Println("Evaluating lambda?")
		namespace := Traj(trajectory, trajectory.Expression)

		for i, param := range trajectory.Expression.Parameters {
			fmt.Println("Defining parameter", param, "as", arguments[i])
			DefineName(&namespace, param, func(t *Trajectory, args ...Value) (Value, error) {
				return arguments[i], nil
			})
		}

		var value Value

		for _, child := range trajectory.Expression.Children {
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

// func EvaluateSet(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	dataV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	data, ok := dataV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	indexV, err := Evaluate(trajectory.Children[1])

// 	index, ok := indexV.(float64)

// 	if !ok {
// 		return nil, errors.New("not a number")
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	val, err := Evaluate(trajectory.Children[2])

// 	if err != nil {
// 		return nil, err
// 	}

// 	data[int(index)] = val

// 	return data, nil
// }

// func EvaluateGet(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	dataV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	data, ok := dataV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	indexV, err := Evaluate(trajectory.Children[1])

// 	index, ok := indexV.(float64)

// 	if !ok {
// 		return nil, errors.New("not a number")
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	val := data[int(index)]

// 	return val, nil
// }

// func EvaluateDo(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	var val Value

// 	for _, input := range trajectory.Children {
// 		v, err := Evaluate(input)

// 		val = v

// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return val, nil
// }

// func EvaluateLogical(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	left, leftErr := Evaluate(trajectory.Children[0])

// 	if leftErr != nil {
// 		return nil, leftErr
// 	}

// 	leftV, ok := left.(bool)

// 	if !ok {
// 		return nil, errors.New("left operand is not a boolean")
// 	}

// 	if trajectory.Expression.Operator.Type == TOKENS.Conjunction {
// 		if !leftV {
// 			return false, nil
// 		}
// 	} else if trajectory.Expression.Operator.Type == TOKENS.Disjunction {
// 		if leftV {
// 			return true, nil
// 		}
// 	} else {
// 		return nil, errors.New("unknown logical operator, && and || are supported")
// 	}

// 	right, rightErr := Evaluate(trajectory.Children[1])

// 	if rightErr != nil {
// 		return nil, rightErr
// 	}

// 	rightV, ok := right.(bool)

// 	if !ok {
// 		return nil, errors.New("right operand is not a boolean")
// 	}

// 	return rightV, nil
// }
