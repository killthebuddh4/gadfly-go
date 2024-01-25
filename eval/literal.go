package eval

import (
	"errors"
	"strconv"
	"strings"

	"github.com/killthebuddh4/gadflai/types"
)

func True(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	return true, nil
}

func False(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	return false, nil
}

func Nil(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	return nil, nil
}

func Number(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	num, parseErr := strconv.ParseFloat(trajectory.Expression.Operator.Value, 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

func String(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	return strings.Trim(trajectory.Expression.Operator.Value, "\""), nil
}
