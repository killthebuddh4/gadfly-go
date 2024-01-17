package array

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Sort(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	compareV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	compare, ok := compareV.(traj.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	sorted := append([]value.Value{}, arr...)

	err = sort(sorted, compare)

	if err != nil {
		return nil, err
	}

	return sorted, nil
}

func sort(arr []value.Value, compare traj.Lambda) error {
	if len(arr) <= 1 {
		return nil
	}

	left, right := 0, len(arr)-1

	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		lessV, err := compare(arr[i], arr[right])

		if err != nil {
			return err
		}

		less, ok := lessV.(float64)

		if !ok {
			return errors.New("less is not a boolean")
		}

		if less < 0 {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}

		arr[left], arr[right] = arr[right], arr[left]

		sort(arr[:left], compare)
		sort(arr[left+1:], compare)
	}

	return nil
}
