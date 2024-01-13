package main

import (
	"errors"
	"strconv"
	"strings"
)

func EvaluateLiteral(trajectory *Trajectory) (Value, error) {
	if trajectory.Expression.Operator.Type == TOKENS.Number {
		return EvaluateNumber(trajectory)
	} else if trajectory.Expression.Operator.Type == TOKENS.String {
		return EvaluateString(trajectory)
	} else {
		return nil, errors.New("unknown literal type <" + trajectory.Expression.Operator.Type + ">")
	}
}

func EvaluateNumber(trajectory *Trajectory) (Value, error) {
	num, parseErr := strconv.ParseFloat(GetLexemeForToken(trajectory.Expression.Operator), 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

func EvaluateString(trajectory *Trajectory) (Value, error) {
	return strings.Trim(trajectory.Expression.Operator.Lexeme, "\""), nil
}

func EvaluateNil(trajectory *Trajectory) (Value, error) {
	return nil, nil
}

func EvaluateTrue(trajectory *Trajectory) (Value, error) {
	return true, nil
}

func EvaluateFalse(trajectory *Trajectory) (Value, error) {
	return false, nil
}
