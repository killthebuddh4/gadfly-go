package exec

import (
	"errors"

	"github.com/killthebuddh4/gadflai/array"
	"github.com/killthebuddh4/gadflai/eval"
	"github.com/killthebuddh4/gadflai/record"
	"github.com/killthebuddh4/gadflai/strings"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

type D func(t *traj.Trajectory, e eval.Eval) (value.Value, error)

func dispatch(trajectory *traj.Trajectory) (D, error) {
	switch trajectory.Expression.Operator.Type {
	case "program":
		return eval.Do, nil
	case "!=":
		return eval.BangEqual, nil
	case "==":
		return eval.EqualEqual, nil
	case ">":
		return eval.Greater, nil
	case ">=":
		return eval.GreaterEqual, nil
	case "<":
		return eval.Less, nil
	case "<=":
		return eval.LessEqual, nil
	case "&&":
		return eval.Conjunction, nil
	case "||":
		return eval.Disjunction, nil
	case "+":
		return eval.Plus, nil
	case "-":
		return eval.Minus, nil
	case "/":
		return eval.Multiply, nil
	case "*":
		return eval.Divide, nil
	case "!":
		return eval.Bang, nil
	case "true":
		return eval.True, nil
	case "false":
		return eval.False, nil
	case "nil":
		return eval.Nil, nil
	case "number":
		return eval.Number, nil
	case "string":
		return eval.String, nil
	case "identifier":
		return eval.Identifier, nil
	case "do":
		return eval.Do, nil
	case "panic":
		return eval.Panic, nil
	case "def":
		return eval.Def, nil
	case "fn":
		return eval.Lambda, nil
	case "@":
		return eval.Call, nil
	case "let":
		return eval.Let, nil
	case "while":
		return eval.While, nil
	case "record":
		return record.Record, nil
	case "merge":
		return record.Merge, nil
	case "delete":
		return record.Delete, nil
	case "extract":
		return record.Extract, nil
	case "keys":
		return record.Keys, nil
	case "values":
		return record.Values, nil
	case "read":
		return record.Read, nil
	case "write":
		return record.Write, nil
	case "array":
		return array.Array, nil
	case "get":
		return array.Get, nil
	case "for":
		return array.For, nil
	case "filter":
		return array.Filter, nil
	case "map":
		return array.Map, nil
	case "reduce":
		return array.Reduce, nil
	case "push":
		return array.Push, nil
	case "pop":
		return array.Pop, nil
	case "set":
		return array.Set, nil
	case "shift":
		return array.Shift, nil
	case "unshift":
		return array.Unshift, nil
	case "segment":
		return array.Segment, nil
	case "find":
		return array.Find, nil
	case "splice":
		return array.Splice, nil
	case "reverse":
		return array.Reverse, nil
	case "sort":
		return array.Sort, nil
	case "if":
		return eval.If, nil
	case "and":
		return eval.And, nil
	case "or":
		return eval.Or, nil
	case "puts":
		return eval.Puts, nil
	case "chars":
		return eval.Chars, nil
	case "split":
		return strings.Split, nil
	case "substring":
		return strings.Substring, nil
	case "concat":
		return strings.Concat, nil
	default:
		return nil, errors.New("error dispatching, unknown operator " + trajectory.Expression.Operator.Type)
	}
}
