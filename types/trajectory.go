package types

import (
	"errors"
	"fmt"
	"os"
)

type Void struct{}

var VOID Void = Void{}

type Thunk func() (Value, error)
type Closure func(context *Trajectory, args ...Value) (Value, error)
type Exec func(scope *Trajectory, args ...Value) (Value, error)

type Trajectory struct {
	Parent      *Trajectory
	Children    []*Trajectory
	Expression  *Expression
	Environment map[string]Value
	Signals     map[string]Closure
	Errors      map[string]Exec
	Arguments   map[string]Value
}

func NewTrajectory(parent *Trajectory, expr *Expression) Trajectory {
	trajectory := Trajectory{
		Parent:      parent,
		Children:    []*Trajectory{},
		Expression:  expr,
		Environment: map[string]Value{},
		Signals:     map[string]Closure{},
		Errors:      map[string]Exec{},
		Arguments:   map[string]Value{},
	}

	return trajectory
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
		_, debug := os.LookupEnv("GADFLY_DEBUG_NAMES")

		if debug {
			fmt.Println("Found name ", key, " while looking for ", name)
		}

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

	_, debug := os.LookupEnv("GADFLY_DEBUG_NAMES")

	if debug {
		fmt.Println("Defining name for ", name)
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

func DefineSignal(trajectory *Trajectory, name string, handler Closure) error {
	if trajectory == nil {
		return errors.New("cannot define signal in nil expression")
	}

	_, ok := trajectory.Signals[name]

	if ok {
		return errors.New("signal " + name + " is already defined")
	}

	trajectory.Signals[name] = handler

	return nil
}

func ResolveSignal(trajectory *Trajectory, name string) (Value, error) {
	if trajectory == nil {
		return nil, errors.New("signal not found for " + name)
	}

	for key, handler := range trajectory.Signals {
		if key == name {
			return handler, nil
		}
	}

	return ResolveSignal(trajectory.Parent, name)
}

func DefineError(trajectory *Trajectory, name string, handler Exec) error {
	if trajectory == nil {
		return errors.New("cannot define error in nil expression")
	}

	_, ok := trajectory.Errors[name]

	if ok {
		return errors.New("error " + name + " is already defined")
	}

	trajectory.Errors[name] = handler

	return nil
}

func ResolveError(trajectory *Trajectory, name string) (Value, error) {
	if trajectory == nil {
		return nil, errors.New("error not found for " + name)
	}

	for key, handler := range trajectory.Errors {
		if key == name {
			return handler, nil
		}
	}

	return ResolveError(trajectory.Parent, name)
}
