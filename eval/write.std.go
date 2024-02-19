package eval

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/killthebuddh4/gadflai/types"
)

var WriteStd types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_EVAL")

	if debug {
		fmt.Println(":: WriteStd :: called")
	}

	for _, arg := range arguments {
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
