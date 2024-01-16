package trajectory

import (
	"errors"

	exp "github.com/killthebuddh4/gadflai/expression"
	"github.com/killthebuddh4/gadflai/value"
)

type Void struct{}

var VOID Void = Void{}

type Lambda func(args ...value.Value) (value.Value, error)

type Trajectory struct {
	Parent      *Trajectory
	Children    []*Trajectory
	Expression  *exp.Expression
	Environment map[string]value.Value
	Yield       value.Value
}

func Traj(parent *Trajectory, expr *exp.Expression) Trajectory {
	return Trajectory{
		Parent:      parent,
		Children:    []*Trajectory{},
		Expression:  expr,
		Environment: map[string]value.Value{},
		Yield:       VOID,
	}
}

func Expand(parent *Trajectory) error {
	children := []*Trajectory{}

	for _, child := range parent.Expression.Children {
		traj := Traj(parent, child)
		children = append(children, &traj)
	}

	parent.Children = children

	return nil
}

func ResolveName(trajectory *Trajectory, name string) (value.Value, error) {
	if trajectory == nil {
		return nil, errors.New("value not found for " + name)
	}

	for key, val := range trajectory.Environment {
		if key == name {
			return val, nil
		}
	}

	return ResolveName(trajectory.Parent, name)
}

func DefineName(trajectory *Trajectory, name string, val value.Value) error {
	if trajectory == nil {
		return errors.New("cannot define name in nil expression")
	}

	_, ok := trajectory.Environment[name]

	if ok {
		return errors.New("name " + name + " is already defined")
	}

	trajectory.Environment[name] = val

	return nil
}

func EditName(trajectory *Trajectory, name string, val value.Value) error {
	if trajectory == nil {
		return errors.New("definition not found for " + name)
	}

	for key := range trajectory.Environment {
		if key == name {
			trajectory.Environment[name] = val
			return nil
		}
	}

	return EditName(trajectory.Parent, name, val)
}
