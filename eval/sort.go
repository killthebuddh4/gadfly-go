package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

type Compare func(types.Value, types.Value) (float64, error)

var Sort types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Sort :: not an array")
	}

	compare, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New("Sort :: not a function")
	}

	sorted := append([]types.Value{}, arr...)

	err := sort(sorted, func(a types.Value, b types.Value) (float64, error) {
		lessV, err := compare(scope, a, b)

		if err != nil {
			return 0, err
		}

		less, ok := lessV.(float64)

		if !ok {
			return 0, errors.New("Sort :: less is not a number")
		}

		return less, nil
	})

	if err != nil {
		return nil, err
	}

	return sorted, nil
}

func sort(arr []types.Value, compare Compare) error {
	if len(arr) <= 1 {
		return nil
	}

	left, right := 0, len(arr)-1

	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		less, err := compare(arr[i], arr[right])

		if err != nil {
			return err
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
