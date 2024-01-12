package main

var BUILTINS = []Definition{
	{
		Name:     "fn",
		Arity:    0,
		Variadic: true,
	},

	{
		Name:     "def",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "val",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "let",
		Arity:    2,
		Variadic: true,
	},

	{
		Name:     "if",
		Arity:    3,
		Variadic: false,
	},

	{
		Name:     "do",
		Arity:    1,
		Variadic: true,
	},

	{
		Name:     "and",
		Arity:    2,
		Variadic: true,
	},

	{
		Name:     "or",
		Arity:    2,
		Variadic: true,
	},

	{
		Name:     "while",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "array",
		Arity:    0,
		Variadic: true,
	},

	{
		Name:     "get",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "set",
		Arity:    3,
		Variadic: false,
	},

	{
		Name:     "for",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "map",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "reduce",
		Arity:    3,
		Variadic: false,
	},

	{
		Name:     "filter",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "push",
		Arity:    2,
		Variadic: false,
	},

	{
		Name:     "pop",
		Arity:    1,
		Variadic: false,
	},

	{
		Name:     "true",
		Arity:    0,
		Variadic: false,
	},

	{
		Name:     "false",
		Arity:    0,
		Variadic: false,
	},

	{
		Name:     "nil",
		Arity:    0,
		Variadic: false,
	},

	{
		Name:     "print",
		Arity:    1,
		Variadic: false,
	},
}
