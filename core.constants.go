package main

func EvaluateTrue(trajectory *Trajectory, args ...Value) (Value, error) {
	return true, nil
}

func EvaluateFalse(trajectory *Trajectory, args ...Value) (Value, error) {
	return false, nil
}

func EvaluateNil(trajectory *Trajectory, args ...Value) (Value, error) {
	return nil, nil
}
