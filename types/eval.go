package types

type Exec func(trajectory *Trajectory) (Value, error)
