package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

// context = caller, scope = the previous child expressions (basically)
func Exec(context *types.Trajectory, scope *types.Trajectory, expr *types.Expression) (types.Value, error) {
	trajectory := types.NewTrajectory(scope, expr)

	if expr.Operator.Type == "fn" && context == scope {
		return close(scope, expr)
	}

	eval, err := dispatch(&trajectory)

	if err != nil {
		return nil, err
	}

	args := []types.Value{}

	for i, param := range expr.Parameters {

		var isThunk bool
		if i >= len(expr.Def.Parameters) {
			isThunk = false
		} else {
			isThunk = expr.Def.Parameters[i].IsThunk
		}

		var arg types.Value

		if isThunk {
			arg = func() (types.Value, error) {
				return Exec(scope, scope, expr)
			}
		} else {
			arg, err = Exec(&trajectory, &trajectory, param)

			if err != nil {
				return nil, err
			}
		}

		args = append(args, arg)
	}

	return eval(&trajectory, args...)
}

func close(scope *types.Trajectory, expr *types.Expression) (types.Closure, error) {
	return func(context *types.Trajectory, arguments ...types.Value) (types.Value, error) {
		injected := types.NewTrajectory(scope, expr)

		if len(arguments) < len(expr.Parameters) {
			return nil, errors.New(":: Exec > close :: not enough arguments")
		}

		// for i, arg := range arguments {
		// 	if i < len(expr.Parameters) {
		// 		types.DefineName(&injected, expr.Parameters[i].Keyword[0].Operator.Value, arg)
		// 	}
		// }

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
	case "then":
		return Then, nil
	case "else":
		return Else, nil
	case "value":
		return Value, nil
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
	case "when":
		return When, nil
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
		return nil, errors.New("error dispatching, unknown operator " + trajectory.Expression.Operator.Type + ".")
	}
}
