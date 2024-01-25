package eval

import (
	"errors"
	"reflect"

	"github.com/killthebuddh4/gadflai/types"
)

func BangEqual(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left != right, nil
}

func EqualEqual(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left == right, nil
}

func Greater(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("Could not evaluate Greater, left operand is not a number. Got " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("Could not evaluate Greater, right operand is not a number. Got " + reflect.TypeOf(right).String())
	}

	return leftV > rightV, nil
}

func GreaterEqual(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("Could not evaluate GreaterEqual, left operand is not a number. Got " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("Could not evaluate GreaterEqual, right operand is not a number. Got " + reflect.TypeOf(right).String())
	}

	return leftV >= rightV, nil
}

func Less(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("Error evaluating Less, left operand is not a number. Got " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("Error evaluating Less, right operand is not a number. Got " + reflect.TypeOf(right).String())
	}

	return leftV < rightV, nil
}

func LessEqual(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("Error evaluating LessEqual, left operand is not a number. Got " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("Error evaluating LessEqual, right operand is not a number. Got " + reflect.TypeOf(right).String())
	}

	return leftV <= rightV, nil
}

func Minus(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	if len(trajectory.Children) == 1 {
		return MinusUnary(trajectory, eval)
	} else {
		return MinusBinary(trajectory, eval)
	}
}

func MinusBinary(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("error evaluating Minus, left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("error evaluating Minus, right operand is not a number")
	}

	return leftV - rightV, nil
}

func MinusUnary(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	right, rightErr := eval(trajectory.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("error evaluating Minus, right operand is not a number")
	}

	return -rightV, nil
}

func Plus(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("error evaluating Plus, left operand is not a number, got " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("error evaluating Plus, right operand is not a number, got " + reflect.TypeOf(right).String())
	}

	return leftV + rightV, nil
}

func Multiply(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("error evaluating Multiply, left operand is not a number, got " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("error evaluating Multiply, right operand is not a number, got " + reflect.TypeOf(right).String())
	}

	return leftV / rightV, nil
}

func Divide(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])
	right, rightErr := eval(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number " + reflect.TypeOf(right).String())
	}

	return leftV * rightV, nil
}

func Bang(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	right, rightErr := eval(trajectory.Children[1])

	if rightErr != nil {
		return nil, errors.New("error evaluating inputs")
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("Error evaluating Bang, right operand is not a boolean. Got " + reflect.TypeOf(right).String())
	}

	return !rightV, nil
}

func Conjunction(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])

	if leftErr != nil {
		return nil, leftErr
	}

	leftV, ok := left.(bool)

	if !ok {
		return nil, errors.New("Error evaluating Conjunction, left operand is not a boolean. Got " + reflect.TypeOf(left).String())
	}

	if !leftV {
		return false, nil
	}

	right, rightErr := eval(trajectory.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("Error evaluating Conjunction, right operand is not a boolean. Got " + reflect.TypeOf(right).String())
	}

	return rightV, nil
}

func Disjunction(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	left, leftErr := eval(trajectory.Children[0])

	if leftErr != nil {
		return nil, leftErr
	}

	leftV, ok := left.(bool)

	if !ok {
		return nil, errors.New("Error evaluating Disjunction, left operand is not a boolean. Got " + reflect.TypeOf(left).String())
	}

	if leftV {
		return true, nil
	}

	right, rightErr := eval(trajectory.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("Error evaluating Disjunction, right operand is not a boolean. Got " + reflect.TypeOf(right).String())
	}

	return rightV, nil
}
