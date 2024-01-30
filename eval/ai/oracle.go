package ai

import (
	"github.com/killthebuddh4/gadflai/types"
)

func Oracle(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	return "oracle", nil
}
