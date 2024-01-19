__Gadfly__ is an experimental expression-oriented functional programming
language and treewalk interpreter. The ultimate goal is to provide a development
environment for [autopoietic](https://en.wikipedia.org/wiki/Autopoiesis) copilot
programs. The conventional aspects of the system are inspired by
[scheme](https://www.scheme.org/) and [ruby](https://www.ruby-lang.org/en/).
The (highly) experimental features draw on ideas from
[cybernetics](https://en.wikipedia.org/wiki/Cybernetics), [symbolic
AI](https://en.wikipedia.org/wiki/Symbolic_artificial_intelligence), and
[reinforcement learning](https://en.wikipedia.org/wiki/Reinforcement_learning)
as well as more recent research surrounding how to build and orchestrate
autonomous agents.

_This project and documentation are under heavy development. If you
see something is missing, find an error, have a question, or have anything at
all to say, please don't hesitate to open an issue or reach out to me directly.
For a (slightly) more detailed overview of the project, check out the [roadmap](#roadmap)_

# Contents

- [Contents](#contents)
- [The language](#the-language)
    - [Blocks](#blocks)
    - [Variables](#variables)
    - [Values](#values)
    - [Lambdas, parameters, and arguments](#lambdas-parameters-and-arguments)
    - [Predicates, operators, and literals](#predicates-operators-and-literals)
    - [Branching](#branching)
    - [Arrays](#arrays)
    - [Maps](#maps)
    - [Strings](#strings)
    - [Input and Output](#input-and-output)
    - [Signals and exceptions](#signals-and-exceptions)
- [Run a script](#run-a-script)
- [Tests](#tests)
- [Notes on the vision](#notes-on-the-vision)
- [Notes on the language](#notes-on-the-language)
- [Notes on the interpreter](#notes-on-the-interpreter)
- [Roadmap](#roadmap)
    - [Phase 1, language core](#phase-1-language-core)
    - [Phase 2, design the cybernetic constructs](#phase-2-design-the-cybernetic-constructs)
    - [Phase 3, intelligence](#phase-3-intelligence)
    - [Nice to haves (unplanned)](#nice-to-haves-unplanned)
- [Work in progress](#work-in-progress)

# The language

Gadfly is dynamically and strongly typed. In Gadfly, everything is a
_lexically-scoped expression_. All expressions return a _value_ and all values
are _immutable_. An __expression__ is defined as a _block_, _lambda_,
_predicate_, or _literal_.  __Comments__ begin with the `#` character and
continue until the end of the line. Whitespace is ignored except to separate
tokens.

For the rest of this section, _`italic text on a gray background`_ denotes an
_expression signature_. For all expression signatures, the `*` character
indicates zero or more occurrences of the preceding expression. The `+`
character indicates one or more occurrences of the preceding expression. The `?`
character indicates an optional expression. Unless otherwise noted, "number",
"string", "array", "map", and "lambda" are understood as expressions that
evaluate to that type of value.

_At the bottom of each subsection in this section you'll find a brief summary of
the ongoing and planned development related to that subsection. For a
higher-level perspective, please see the [roadmap](#roadmap)._

### Blocks

A __block__ is a sequence of expressions delimited by a _keyword_ and `end`. A
__keyword__ determines its block's behavior or semantics. Most of the language's
keywords will be described throughout the rest of this section but you can also
find a comprehensive, runnable example in
[examples.core.fly](examples.core.fly).

The simplest block is the `do` block:

_`do expression* end`_

The expressions are evaluated in order and the value of the last expression is
returned.

```text
do
  puts "hey" end

  2

  do
    3 + 4
  end
end
```

### Variables

A Gadfly __variable__ is an expression that resolves to a _value_ by referencing
it. A variable is defined using a `def` block and re-defined using a `let`
block. After a variable is defined it can be referenced in any expression.

_`def identifier expression end`_

Defines a variable with the given identifier. The variable resolves to the value
of the expression. Variables are _lexically scoped_. If the variable is already
defined in the local scope, it is an error. If the variable is defined in an
outer scope, it will be _shadowed_ in the local scope.

```text
def surname "smith" end
```

_`let identifier expression end`_

Re-defines an existing variable with the given identifier. The variable resolves
to the value of the expression. If the variable does not already exist, it is
an error.

```
def val "hi" end
let val "goodbye" end
```

__TODO__

- [ ] Namespace declaration and resolution.

### Values

Every value is a _string_, _number_, _array_, _map_, _lambda_, or _nil_.

A __string__ is created by enclosing characters in quotes.

```text
"I am string"
```

A __number__ is created by writing it out in decimal notation. All numbers are
represented as floats internally.

```text
1
0.1
10.0
```

There is no _boolean_ type in Gadfly. All "boolean" operators take _number_
operands and treat `0` as false-y and any other number as truth-y. All other
values cause errors when used as a boolean.

An __array__ is created using the `array` block and is a number-indexed list of
values. See the [arrays](#arrays) section for more details on arrays.

A __map__ is created using the `map` block and is a string-keyed dictionary of
values. See the [maps](#maps) section for more details on maps.

A __lambda__ is created using the `fn` block and can be thought of as a
parameterized _do_ block or "anonymous function". See the [lambdas](#lambdas)
section for more details on lambdas.

### Lambdas, parameters, and arguments

A __lambda__ is a "parameterized block" that is not evaluated until each time it
is called. A lambda can have zero or more _parameters_. A __parameter__ is a
name that is defined each time the lambda is called. Parameters are declared
between `|` characters. If the lambda takes zero parameters, the `|` characters
must be omitted. The  __arguments__ to the lambda are the values of the
expressions in the calling block (using the `@` keyword) bound to the lambda's
parameters.

_`fn (|identifier+|)? expression end`_

When the lambda expression is evaluated, it creates a lambda. The key difference
between a lambda expression and other expressions is that its subexpressions are
evaluated only when the lambda is called. The lambda can take zero or more
parameters. If the lambda takes zero parameters, the `|` characters must be
omitted.

`@ expression* end`

Calls the lambda expression. Each subexpression is evaluated and bound to the
lambda's parameters. The lambda is then evaluated, returning the value of its
last subexpression.

```text
def add
  # parameters are a and b
  fn |a b|
    a + b
  end
end

@add
  # arguments are 8 and 3, bound to a and b
  2 * 4
  3
end

map
  array 1 2 3 end

  fn |n i|
    n + i
  end
end
```

### Predicates, operators, and literals

A __predicate__ is an expression involving an _operator_ and _operands_. See the
[operators](#operators) section for more details on each operator. An
__operand__ is either a _predicate_ or a _literal_. A __literal__ is an
expression without subexpressions (string, number, boolean, variable). A
predicate evaluates to a _number_ (because an operator evaluates to a number).

_Because predicates cannot include blocks they cannot include function calls.
This is somewhat cumbersome to us human programmers, forcing us to write many
instances of trivial indirection, but I think we'll see strong benefits for code
generation and program synthesis because it will make parse trees simpler. Maybe
not, we'll see._

```text
# Not predicates.

fn
  io.puts "hi" end
end

def val "hi" end

# Predicates.

val

val == "goodbye"

10 > 0 # => 1

100 / 20 # => 5

!val
```

__TODO__

- [ ] Unary negation seems to be broken right now.

### Branching

The key difference between branching expressions and other expressions is that
their subexpression are evaluated conditionally. The specific behavior of which
subexpressions are evaluated depends on the keyword.

_Note that branching expressions are not predicates, they may return any value._

_`if number expression expression end`_

If the number is truth-y, the first expression is evaluated. Otherwise, the
second expression is evaluated. The value of the last evaluated expression is
returned.


_`and (number expression)+ end`_

For each pair of subexpressions, if the first evaluates to a truth-y value, the
second is evaluated. If any of the subexpressions evaluate to a false-y value,
`nil` is returned. Otherwise, the value of the last subexpression is returned.

_`or (number expression)+ end`_

For each pair of subexpressions, if the first evaluates to a truth-y value, the
second is evaluated and returned. If none of the subexpressions evaluate to a
truth-y value, `nil` is returned.

_`while number expression+ end`_

While the first expression evaluates to a truth-y value, the rest of the expressions
are evaluated. The value of the last subexpression is returned.

### Arrays

_`array expression* end`_

Creates an array whose values are the values of the subexpressions. The array is 
returned.

_`array.read array number end`_

The value of the array at the index of the number is returned.

_`array.write array number expression end`_

Clones the array and sets the value at the index of the number to the value of
the expression. The cloned array is returned.

_`array.for array lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
The value of the last evaluated lambda is returned.

_`array.map array lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
An array whose values are the result of each lambda call is returned.

_`array.filter array lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
An array whose values are the values for which the lambda call returned a
truth-y value is returned.

_`array.reduce array expression lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's second parameter and the index bound to the lambda's third parameter.
When the lambda is called for the first value in the array, the first parameter
is bound to the value of expression. For each subsequent value in the array, the
first parameter is bound to the value returned by the previous lambda call. The
value of the last evaluated lambda is returned.

_`array.push array expression end`_

Clones the array and appends the value of the expression to the cloned array. 
The cloned array is returned.

_`array.pop array end`_

Clones the array and removes the last value from the cloned array. The cloned
array is returned.

_`array.unshift array expression end`_

Clones the array and prepends the value of the expression to the cloned array.
The cloned array is returned.

_`array.shift array end`_

Clones the array and removes the first value from the cloned array. The cloned
array is returned.

_`array.reverse array end`_

Clones the array and reverses the order of the values in the cloned array. The
cloned array is returned.

_`array.sort array lambda end`_

Clones the array and sorts the values in the cloned array according to the value
returned by the lambda. The lambda takes two parameters, the values of which are
the values in the array. The lambda returns a negative number if the first value
should be sorted before the second, a positive number if the first value should
be sorted after the second, and `0` if the values are equal. The cloned (sorted)
array is returned.

_`array.segment array number number end`_

Clones the array and returns a new array whose values are the values of the
cloned array between the first index and the second index (exclusive). The
cloned array is returned.

_`array.splice array number array end`_

Clones the first array and divides it in half at the index of the number. It
appends the values of the second array to the first half, and then appends the
second half to the result. The result is returned.

### Maps

_`map (string expression)* end`_

Creates a map whose keys are the strings and whose values are the values of
the expressions. The map is returned.

_`map.read map string end`_

The value of the map at the key of the string is returned.

_`map.write map string expression end`_

Clones the map and sets the value at the key of the string to the value of
the expression. The cloned map is returned.

_`map.delete map array end`_

The array is an array of strings. Clones the map and deletes the keys of the
strings from the cloned map. The cloned map is returned.

_`map.extract map array end`_

The array is an array of strings. Returns a map whose keys are the keys of
the strings and whose values are the values of the keys of the strings in the
map. The new map is returned.

_`map.merge map map end`_

Clones the first map and then for each kv pair in the second map, sets the
value of the cloned map at the key of the kv pair to the value of the kv
pair. Returns the cloned map.

_`map.keys map end`_

An array whose values are the keys of the map is returned.

_`map.values map end`_

An array whose values are the values of the map is returned.

### Strings

_`split string end`_

Returns an array whose values are the characters in the string.

_`concat string+ end`_

Returns a string whose value is the concatenation of the values of the strings.

_`substring string number number end`_

Returns a string whose value is the substring of the string between the first
index and the second index (exclusive).

__TODO__

- [ ] regular expression engine

### Input and Output

__TODO__

- [x] `io.puts`
- [ ] `io.gets`
- [ ] `io.err`
- [ ] _WIP_ `io.http`
- [ ] `io.file.read`
- [ ] `io.file.write`
- [ ] documentation

### Signals and exceptions

__TODO__

- [ ] A Lisp-style _condition_ system but more focused on _signaling_.
- [ ] documentation

# Run a script

_Requires `go` 1.21 or higher. Learn how to install `go` [here](https://go.dev/doc/install)._

```bash
go run . <path to Gadfly source>
```

Try running the examples:

```bash
for file in examples/*.fly; do
  go run . $file
done
```

# Tests

 The goal is to have tests commensurate with the maturity of the project and its
 components. The near term goal is to have something like 100% coverage for the
 core language features. Basically, this means "all of the keywords and
 operators". We'll do this incrementally, in phases.

 You can run the tests with:

 ```bash
 ./test.sh
 ```

__TODO__

 - [ ] _WIP_ Smoke coverage for all keywords and operators
  - [x] array
  - [x] strings
  - [ ] map
  - [ ] branching
  - [ ] lambdas
  - [ ] io
    - [x] http
  - [ ] namespaces
  - [ ] emitters
  - [ ] exceptions
  - [ ] variables
  - [ ] predicates
 - [ ] Edge-case coverage for all keywords and operators
 - [ ] Happy-path coverage for at least one robust, traditional Gadfly program.
 - [ ] _Down the road_ Fuzzing for all keywords and operators

# Notes on the vision

- Imagine a programming language designed from the ground up to be used by
  language models (LMs).
- An LM wouldn't need to generalize, it could write a new program for each new
  task it encounters.
- An LM could analyze running programs very quickly, it could modify
  running programs according to new data.
- Instead of importing code, an LM could search for similar code in a database
  and then repurpose it.
- Because an LM could analyze running programs very quickly, it could make
  sense to store all kinds of useful metadata in the parse tree.
- More generally, the parse tree could be something that is constantly
  manipulated by the LM.
- An LM could outsource subtrees to other LMs, there could be a whole ecosystem
  built around this idea.
- An LM could evaluate subtrees in parallel and then merge, prune, and recombine
  them according to certain policies.

# Notes on the language

- It's looking like the language will (unsurprisingly) be very Lispy. One way to
  think about things is that `Gadfly` takes homoiconicity to wild extremes.
- Right now the language is dynamically typed, but I feel like it would be
  better to have static typing. I don't have any great reasons to articulate
  why I feel that way other than maybe that it adds a ton of metadata to the
  parse tree. I don't know enough about programming languages to know whether
  lisps are amenable to static typing.
- I have this vague notion of a `remote` keyword, or something like that, where
  a subtree is farmed out to the network. I feel like a `remote` node could be
  implemented in any language and basically implement any feature.
- The syntax is intended to mirror the parse tree very closely. I imagine that
  developers will jump around _a lot_ between the two. I want the
  _homoiconicity_ to be very apparent.
- Everything about the language needs to be ridiculously printable,
  introspectable, and serializable. I want to be able to seralize the parse
  tree, wire it somewhere, and then execute it on another machine.

# Notes on the interpreter

- you have a parse tree as well as metadata surrounding every execution that has
  ever occurred in the parse tree
- you never edit a node in the parse tree, you only deprecate them
- the parse tree represents a decision making process that the AI will incorporate
- one way to think about how this works is that it uses LLMs and execution
  history to discretize the continuous decision spac
- one way to think about what the copilot is doing is “how to write a program
  that proves a fact about the world”
- similar to the previous bullet, there’s a way to think about is synthesizing
  analytical facts and a fact is a subtree that can prove the thing
- the copilot’s toolbox is determined via lexically-scoped name resolution
- instead of a visual programming environment, you have a programming language
  which means you can use programming language theory and tools
- a core aspect of a copilot is that it must support delegation to other AIs
- every expression in the parse tree has a natural language representation of
  why it exists in the parse tree
- every node in a trajectory has a natural language description of it’s scope,
  it’s arguments, and a recapitulation
- one piece of metadata is “number of active trajectories”
- the parse tree must be fully serializable
- you can think of every subtree as a policy fit for some situation
- a gadflai program has something like sensors or actuators
- leaf nodes are always either facts or sensors?
- a pure function is a fact? something with an effect is a sensor?
- every subtree can have a supervisor? or every subtree can have a policy?

# Roadmap

### Phase 1, language core

- [x] design and implement the architecture
  - [x] lex
  - [x] parse
  - [x] eval
  - [x] closures
- [ ] mvp language features
  - [x] array
  - [x] strings
  - [x] map
  - [x] branching
  - [x] lambdas
  - [x] variables
  - [x] predicates
  - [ ] _WIP_ io
  - [ ] namespaces
  - [ ] emitters
  - [ ] signals
- [ ] syntax highlighting

### Phase 2, design the cybernetic constructs

_Like an exosuit for language models?_

- [ ] the copilot architecture (analysis, synthesis, proving, observation, etc.)
- [ ] the user flow
- [ ] the persistence layer
- [ ] feedback and expansion
- [ ] delegation and orchestration
- [ ] Implement MVP tooling for the above and for the next phase

### Phase 3, intelligence

- [ ] Prompt generation
- [ ] observation mechanisms
- [ ] evaluation mechanisms
- [ ] retry mechanisms
- [ ] shuffling mechanisms
- [ ] delegation and orchestration

### Nice to haves (unplanned)

- [ ] language server protocol implementation
- [ ] repl
- [ ] tail call optimization
- [ ] static typing

# Work in progress

- [ ] add an intepreter section at the top of the README
- [ ] merge the syntax and semantics sections in the README
- [ ] http keyword
- [ ] namespaces
- [ ] emitters
- [ ] try to get a rough draft of the interpreter section using notes