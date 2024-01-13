package main

import "errors"

func EvaluateChars(scope *Trajectory) (Value, error) {
	expand(scope)

	arg, err := evaluate(scope.Children[0])

	if err != nil {
		return nil, err
	}

	str, strOk := arg.(string)

	if !strOk {
		return nil, errors.New("chars only accepts strings")
	}

	result := []Value{}

	for _, c := range str {
		result = append(result, string(c))
	}

	return result, nil
}
