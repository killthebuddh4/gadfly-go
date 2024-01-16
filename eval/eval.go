package eval

import (
	"github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

type Eval func(trajectory *trajectory.Trajectory) (value.Value, error)
