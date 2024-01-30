package types

import (
	"errors"
)

type Void struct{}

var VOID Void = Void{}

type Lambda func(args ...Value) (Value, error)

type Trajectory struct {
	Parent      *Trajectory
	Children    []*Trajectory
	Expression  *Expression
	Environment map[string]Value
	Yield       Value
}

func NewTrajectory(parent *Trajectory, expr *Expression) Trajectory {
	return Trajectory{
		Parent:      parent,
		Children:    []*Trajectory{},
		Expression:  expr,
		Environment: map[string]Value{},
		Yield:       VOID,
	}
}

func ExpandBy(parent *Trajectory, exp *Expression) error {
	var isChildExpression bool = false

	for _, child := range parent.Expression.Children {
		if child == exp {
			isChildExpression = true
		}
	}

	if !isChildExpression {
		return errors.New("expression is not a child of parent")
	}

	var isAlreadyExpanded bool = false

	for _, child := range parent.Children {
		if child.Expression == exp {
			isAlreadyExpanded = true
		}
	}

	if isAlreadyExpanded {
		return errors.New("expression is already expanded")
	}

	child := NewTrajectory(parent, exp)
	parent.Children = append(parent.Children, &child)

	return nil
}

func ExpandTraj(parent *Trajectory) error {
	children := []*Trajectory{}

	for _, child := range parent.Expression.Children {
		child := NewTrajectory(parent, child)
		children = append(children, &child)
	}

	parent.Children = children

	return nil
}

func ResolveName(trajectory *Trajectory, name string) (Value, error) {
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

func DefineName(trajectory *Trajectory, name string, val Value) error {
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

func EditName(trajectory *Trajectory, name string, val Value) error {
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
