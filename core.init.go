package main

func InitStdlib(root *Trajectory) error {
	err := DefineName(root, "print", EvaluatePrint)

	if err != nil {
		return err
	}

	err = DefineName(root, "chars", EvaluateChars)

	if err != nil {
		return err
	}

	err = DefineName(root, "def", EvaluateDef)

	if err != nil {
		return err
	}

	err = DefineName(root, "fn", EvaluateLambda)

	if err != nil {
		return err
	}

	err = DefineName(root, "do", EvaluateDo)

	if err != nil {
		return err
	}

	err = DefineName(root, "if", EvaluateIf)

	if err != nil {
		return err
	}

	err = DefineName(root, "and", EvaluateAnd)

	if err != nil {
		return err
	}

	err = DefineName(root, "or", EvaluateOr)

	if err != nil {
		return err
	}

	err = DefineName(root, "for", EvaluateFor)

	if err != nil {
		return err
	}

	return nil
}
