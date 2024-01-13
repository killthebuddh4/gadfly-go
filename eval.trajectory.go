package main

import (
	"errors"
	"fmt"
)

type Void struct{}

var VOID Void = Void{}

type Trajectory struct {
	Parent      *Trajectory
	Children    []*Trajectory
	Expression  *Expression
	Environment map[string]Lambda
	Yield       Value
}

func Traj(parent *Trajectory, expr *Expression) Trajectory {
	return Trajectory{
		Parent:      parent,
		Children:    []*Trajectory{},
		Expression:  expr,
		Environment: make(map[string]Lambda),
		Yield:       VOID,
	}
}

func expand(parent *Trajectory) error {
	children := []*Trajectory{}

	for _, child := range parent.Expression.Children {
		traj := Traj(parent, child)
		children = append(children, &traj)
	}

	parent.Children = children

	fmt.Println("Expanded <", parent.Expression.Operator.Lexeme, "> to N <", len(parent.Children), ">")

	return nil
}

func ResolveName(trajectory *Trajectory, name string) (Lambda, error) {
	if trajectory == nil {
		return nil, errors.New("value not found for <" + name + ">")
	}

	for key, val := range trajectory.Environment {
		if key == name {
			return val, nil
		}
	}

	return ResolveName(trajectory.Parent, name)
}

func DefineName(trajectory *Trajectory, name string, val Lambda) error {
	if trajectory == nil {
		return errors.New("cannot define name in nil expression")
	}

	_, ok := trajectory.Environment[name]

	if ok {
		return errors.New("name " + name + " is already defined")
	}

	fmt.Println("Defining name <", name, "> in trajectory <", trajectory.Expression.Operator.Type, ">")
	trajectory.Environment[name] = val

	return nil
}

func EditName(trajectory *Trajectory, name string, val Lambda) error {
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
