package main

import "errors"

// func EvaluateWhile(trajectory *Trajectory, args ...Value) (Value, error) {
// 	expand(trajectory)

// 	var value Value = nil

// 	for {
// 		condV, err := Evaluate(trajectory.Children[0])

// 		if err != nil {
// 			return nil, err
// 		}

// 		cond, ok := condV.(bool)

// 		if !ok {
// 			return nil, errors.New("not a boolean")
// 		}

// 		if !cond {
// 			break
// 		} else {
// 			for _, child := range trajectory.Children[1:] {
// 				val, err := Evaluate(child)

// 				if err != nil {
// 					return nil, err
// 				}

// 				value = val
// 			}
// 		}
// 	}

// 	return value, nil
// }

func EvaluateFor(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	arrV, err := Evaluate(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	for i, v := range arr {
		scopeExp := Expr(
			nil,
			VARIANTS.Call,
			Token{
				Type:   TOKENS.Identifier,
				Start:  -1,
				Length: -1,
				Lexeme: "pseudo",
			},
		)

		Expr(scopeExp, VARIANTS.Literal, Token{
			Type:   TOKENS.Number,
			Start:  -1,
			Length: -1,
			Lexeme: "0",
		})

		Expr(scopeExp, VARIANTS.Literal, Token{
			Type:   TOKENS.Number,
			Start:  -1,
			Length: -1,
			Lexeme: "0",
		})

		traj := Traj(trajectory, scopeExp)
		_, err := fn(nil)

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// func EvaluateMap(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	arrV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	fnV, err := Evaluate(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fn, ok := fnV.(Lambda)

// 	if !ok {
// 		return nil, errors.New("not a function")
// 	}

// 	vals := []Value{}

// 	for i, v := range arr {
// 		mapped, err := fn(v, float64(i))

// 		if err != nil {
// 			return nil, err
// 		}

// 		vals = append(vals, mapped)
// 	}

// 	return vals, nil
// }

// func EvaluatePush(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	arrV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	val, err := Evaluate(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr = append(arr, val)

// 	return arr, nil
// }

// func EvaluatePop(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	arrV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	return arr[:len(arr)-1], nil
// }

// func EvaluateFilter(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	arrV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	fnV, err := Evaluate(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fn, ok := fnV.(Lambda)

// 	if !ok {
// 		return nil, errors.New("not a function")
// 	}

// 	vals := []Value{}

// 	for i, v := range arr {
// 		filterV, err := fn(v, float64(i))

// 		if err != nil {
// 			return nil, err
// 		}

// 		filter, ok := filterV.(bool)

// 		if !ok {
// 			return nil, errors.New("filter is not a boolean")
// 		}

// 		if filter {
// 			vals = append(vals, v)
// 		}
// 	}

// 	return vals, nil
// }

// func EvaluateReduce(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	arrV, err := Evaluate(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	initV, err := Evaluate(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fnV, err := Evaluate(trajectory.Children[2])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fn, ok := fnV.(Lambda)

// 	if !ok {
// 		return nil, errors.New("not a function")
// 	}

// 	if (len(arr)) == 0 {
// 		return nil, nil
// 	}

// 	reduction := initV

// 	for i, v := range arr {
// 		reduction, err = fn(reduction, v, float64(i))

// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return reduction, nil
// }

// func EvaluateArray(trajectory *Trajectory) (Value, error) {
// 	expand(trajectory)

// 	arr := []Value{}

// 	for _, input := range trajectory.Children {
// 		val, err := Evaluate(input)

// 		if err != nil {
// 			return nil, err
// 		}

// 		arr = append(arr, val)
// 	}

// 	return arr, nil
// }
