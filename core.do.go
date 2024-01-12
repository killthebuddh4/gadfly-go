package main

func EvaluateDo(scope *Trajectory, args ...Value) (Value, error) {
	if len(args) == 0 {
		return nil, nil
	}

	return args[len(args)-1], nil
}
