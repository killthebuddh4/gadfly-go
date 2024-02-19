package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Exec(context *types.Trajectory, scope *types.Trajectory, expr *types.Expression) (types.Value, error) {
	trajectory := types.NewTrajectory(scope, expr)

	if expr.Operator.Type == "fn" && context == scope {
		return close(scope, expr)
	} else {
		eval, err := dispatch(&trajectory)

		if err != nil {
			return nil, err
		}

		args := []types.Value{}

		for _, child := range expr.Children {
			value, err := Exec(&trajectory, &trajectory, child)

			if err != nil {
				return nil, err
			}

			args = append(args, value)
		}

		return eval(&trajectory, args...)
	}
}

func close(scope *types.Trajectory, expr *types.Expression) (types.Closure, error) {
	return func(context *types.Trajectory, arguments ...types.Value) (types.Value, error) {
		injected := types.NewTrajectory(scope, expr)

		if len(arguments) < len(expr.Parameters) {
			return nil, errors.New(":: Exec > close :: not enough arguments")
		}

		for i, arg := range arguments {
			if i < len(expr.Parameters) {
				types.DefineName(&injected, expr.Parameters[i].Children[0].Operator.Value, arg)
			}
		}

		return Exec(context, &injected, expr)
	}, nil
}

func dispatch(trajectory *types.Trajectory) (types.Exec, error) {
	switch trajectory.Expression.Operator.Type {
	case "program":
		return Do, nil
	case "!=":
		return BangEqual, nil
	case "==":
		return EqualEqual, nil
	case ">":
		return Greater, nil
	case ">=":
		return GreaterEqual, nil
	case "<":
		return Less, nil
	case "<=":
		return LessEqual, nil
	case "&&":
		return Conjunction, nil
	case "||":
		return Disjunction, nil
	case "+":
		return Plus, nil
	case "-":
		return Minus, nil
	case "/":
		return Multiply, nil
	case "*":
		return Divide, nil
	case "!":
		return Bang, nil
	case "true":
		return True, nil
	case "false":
		return False, nil
	case "nil":
		return Nil, nil
	case "symbol":
		return Symbol, nil
	case "number":
		return Number, nil
	case "string":
		return String, nil
	case "identifier":
		return Identifier, nil
	case "do", "fn":
		return Do, nil
	case "panic":
		return Panic, nil
	case "def":
		return Def, nil
	case ".":
		return Call, nil
	case "let":
		return Let, nil
	case "while":
		return While, nil
	case "map":
		return Hash, nil
	case "map.merge":
		return Merge, nil
	case "map.delete":
		return Delete, nil
	case "map.extract":
		return Extract, nil
	case "map.keys":
		return Keys, nil
	case "map.values":
		return Values, nil
	case "map.read":
		return ReadHash, nil
	case "map.write":
		return WriteHash, nil
	case "array":
		return Array, nil
	case "array.for":
		return For, nil
	case "array.filter":
		return Filter, nil
	case "array.map":
		return Map, nil
	case "array.reduce":
		return Reduce, nil
	case "array.push":
		return Push, nil
	case "array.pop":
		return Pop, nil
	case "array.read":
		return ReadArray, nil
	case "array.write":
		return WriteArray, nil
	case "array.shift":
		return Shift, nil
	case "array.unshift":
		return Unshift, nil
	case "array.segment":
		return Segment, nil
	case "array.find":
		return Find, nil
	case "array.splice":
		return Splice, nil
	case "array.reverse":
		return Reverse, nil
	case "array.sort":
		return Sort, nil
	case "if":
		return If, nil
	case "and":
		return And, nil
	case "or":
		return Or, nil
	case "case":
		return Case, nil
	case "std.write":
		return WriteStd, nil
	case "http":
		return Http, nil
	case "chars":
		return Chars, nil
	case "split":
		return Split, nil
	case "substring":
		return Substring, nil
	case "concat":
		return Concat, nil
	case "GADFLY":
		return Gadfly, nil
	case "DAEMON":
		return Daemon, nil
	case "GHOST":
		return Ghost, nil
	case "ORACLE":
		return Oracle, nil
	case "MUSE":
		return Muse, nil
	case "RAPTURE":
		return Rapture, nil
	case ":":
		return Colon, nil
	case "schema":
		return Schema, nil
	case "catch":
		return Catch, nil
	case "throw":
		return Throw, nil
	case "signal":
		return Signal, nil
	case "emit":
		return Emit, nil
	case "on":
		return On, nil
	default:
		return nil, errors.New("error dispatching, unknown operator " + trajectory.Expression.Operator.Type)
	}
}

// func Exec(trajectory *types.Trajectory) (types.Value, error) {
// 	eval, dispatchErr := dispatch(trajectory)

// 	if dispatchErr != nil {
// 		return nil, dispatchErr
// 	}

// 	if trajectory.Expression.Operator.Value != "fn" {
// 		for _, child := range trajectory.Expression.Parameters {
// 			validationTrajectory := types.NewTrajectory(trajectory, child)

// 			eval, dispatchErr := dispatch(&validationTrajectory)

// 			if dispatchErr != nil {
// 				return nil, dispatchErr
// 			}

// 			_, err := eval(&validationTrajectory, Exec)

// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	val, evalErr := eval(trajectory, Exec)

// 	if evalErr != nil {
// 		return nil, evalErr
// 	}

// 	if trajectory.Expression.Operator.Value == "fn" {
// 		return val, nil
// 	} else {
// 		if len(trajectory.Expression.Returns) == 0 {
// 			return val, nil
// 		} else {
// 			validationTrajectory := types.NewTrajectory(trajectory, trajectory.Expression.Returns[0])
// 			eval, dispatchErr := dispatch(&validationTrajectory)

// 			if dispatchErr != nil {
// 				return nil, dispatchErr
// 			}

// 			schemaV, err := eval(&validationTrajectory, Exec)

// 			if err != nil {
// 				return nil, err
// 			}

// 			schema, ok := schemaV.(types.Lambda)

// 			if !ok {
// 				return nil, fmt.Errorf("not a function")
// 			}

// 			return schema(val)
// 		}
// 	}
// }
