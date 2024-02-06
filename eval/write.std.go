package eval

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/killthebuddh4/gadflai/types"
)

func WriteStd(trajecotry *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajecotry)

	args := []types.Value{}

	for _, child := range trajecotry.Children {
		arg, err := eval(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	for _, arg := range args {
		m, mOk := arg.(map[string]types.Value)
		str, strOk := arg.(string)
		float, floatOk := arg.(float64)
		i, intOk := arg.(int)
		tf, tfOk := arg.(bool)
		slice, sliceOk := arg.([]types.Value)

		if arg == nil {
			fmt.Println("nil")
		} else if mOk {
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
			return nil, errors.New("std.write only accepts booleans and strings and numbers, got " + reflect.TypeOf(arg).String())
		}
	}

	return nil, nil
}
