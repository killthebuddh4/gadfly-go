package exec

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
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
	case "array":
		return eval.Array, nil
	case "set":
		return eval.Set, nil
	case "get":
		return eval.Get, nil
	case "for":
		return eval.For, nil
	case "filter":
		return eval.Filter, nil
	case "map":
		return eval.Map, nil
	case "reduce":
		return eval.Reduce, nil
	case "push":
		return eval.Push, nil
	case "pop":
		return eval.Pop, nil
	case "shift":
		return eval.Shift, nil
	case "unshift":
		return eval.Unshift, nil
	case "segment":
		return eval.Segment, nil
	case "find":
		return eval.Find, nil
	case "splice":
		return eval.Splice, nil
	case "reverse":
		return eval.Reverse, nil
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
	default:
		return nil, errors.New("error dispatching, unknown operator " + trajectory.Expression.Operator.Type)
	}
}
