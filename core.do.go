package main

func EvaluateDo(scope *Trajectory) (Value, error) {
	expand(scope)

	var value Value

	for _, child := range scope.Children {
		val, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		value = val
	}

	return value, nil
}
