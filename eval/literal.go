package eval

import (
	"errors"
	"strconv"
	"strings"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func True(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	return true, nil
}

func False(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	return false, nil
}

func Nil(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	return nil, nil
}

func Number(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	num, parseErr := strconv.ParseFloat(trajectory.Expression.Operator.Value, 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

func String(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	return strings.Trim(trajectory.Expression.Operator.Value, "\""), nil
}
