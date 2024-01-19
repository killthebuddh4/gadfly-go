package io

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Puts(trajecotry *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajecotry)

	args := []value.Value{}

	for _, child := range trajecotry.Children {
		arg, err := eval(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	for _, arg := range args {
		m, mOk := arg.(map[string]value.Value)
		str, strOk := arg.(string)
		float, floatOk := arg.(float64)
		i, intOk := arg.(int)
		tf, tfOk := arg.(bool)
		slice, sliceOk := arg.([]value.Value)

		if arg == nil {
			fmt.Println("nil")
		} else if mOk {
			fmt.Println("record")
			for k, v := range m {
				fmt.Printf("    %s: ", k)
				fmt.Println(v)
			}
			fmt.Println("end")
		} else if strOk {
			fmt.Println(str)
		} else if floatOk {
			fmt.Println(float)
		} else if intOk {
			fmt.Println(i)
		} else if tfOk {
			fmt.Println(tf)
		} else if sliceOk {
			fmt.Println("array")
			for _, v := range slice {
				fmt.Printf("    ")
				fmt.Println(v)
			}
			fmt.Println("end")
		} else {
			return nil, errors.New("io.puts only accepts booleans and strings and numbers, got " + reflect.TypeOf(arg).String())
		}
	}

	return nil, nil
}
