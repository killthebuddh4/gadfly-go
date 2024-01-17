package eval

// import (
// 	"errors"
// 	"fmt"

// 	traj "github.com/killthebuddh4/gadflai/trajectory"
// 	"github.com/killthebuddh4/gadflai/value"
// )

// func Merge(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	panic("not implemented")

// 	traj.Expand(trajectory)

// 	dataV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	data, ok := dataV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("error setting array, data is not an array, it is " + fmt.Sprint(dataV))
// 	}

// 	indexV, err := eval(trajectory.Children[1])

// 	index, ok := indexV.(float64)

// 	if !ok {
// 		return nil, errors.New("error setting array, index is not a number, it is " + fmt.Sprint(indexV))
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	val, err := eval(trajectory.Children[2])

// 	if err != nil {
// 		return nil, err
// 	}

// 	data[int(index)] = val

// 	return data, nil
// }

// func Get(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	dataV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	data, ok := dataV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("error getting from array, data is not an array, it is " + fmt.Sprint(dataV))
// 	}

// 	indexV, err := eval(trajectory.Children[1])

// 	index, ok := indexV.(float64)

// 	if !ok {
// 		return nil, errors.New("error getting from array, index is not a number, it is " + fmt.Sprint(indexV))
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	val := data[int(index)]

// 	return val, nil
// }

// func Map(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	fnV, err := eval(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fn, ok := fnV.(traj.Lambda)

// 	if !ok {
// 		return nil, errors.New("not a function")
// 	}

// 	vals := []value.Value{}

// 	for i, v := range arr {
// 		mapped, err := fn(v, float64(i))

// 		if err != nil {
// 			return nil, err
// 		}

// 		vals = append(vals, mapped)
// 	}

// 	return vals, nil
// }

// func Push(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	val, err := eval(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr = append(arr, val)

// 	return arr, nil
// }

// func Pop(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	return arr[:len(arr)-1], nil
// }

// func Shift(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	panic("not implemented")

// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	val, err := eval(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr = append(arr, val)

// 	return arr, nil
// }

// func Unshift(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	panic("not implemented")

// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	return arr[:len(arr)-1], nil
// }

// func Filter(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	fnV, err := eval(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fn, ok := fnV.(traj.Lambda)

// 	if !ok {
// 		return nil, errors.New("not a function")
// 	}

// 	vals := []value.Value{}

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

// func Reduce(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	arrV, err := eval(trajectory.Children[0])

// 	if err != nil {
// 		return nil, err
// 	}

// 	arr, ok := arrV.([]value.Value)

// 	if !ok {
// 		return nil, errors.New("not an array")
// 	}

// 	initV, err := eval(trajectory.Children[1])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fnV, err := eval(trajectory.Children[2])

// 	if err != nil {
// 		return nil, err
// 	}

// 	fn, ok := fnV.(traj.Lambda)

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

// func Array(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
// 	traj.Expand(trajectory)

// 	arr := []value.Value{}

// 	for _, input := range trajectory.Children {
// 		val, err := eval(input)

// 		if err != nil {
// 			return nil, err
// 		}

// 		arr = append(arr, val)
// 	}

// 	return arr, nil
// }
