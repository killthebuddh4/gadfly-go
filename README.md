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

_Please note, this project and documentation are under heavy development. If you
see something is missing, find an error, have a question, or have anything at
all to say, an [issue](https://github.com/killthebuddh4/gadfl.ai/issues) would
be awesome!_

# Contents

- [Contents](#contents)
- [Language core and syntax](#language-core-and-syntax)
    - [Expressions](#expressions)
    - [Lambdas, parameters, and arguments](#lambdas-parameters-and-arguments)
    - [Predicates, operators, and literals](#predicates-operators-and-literals)
    - [Variables](#variables)
    - [Values](#values)
- [Semantics](#semantics)
    - [Variables](#variables-1)
    - [Lambdas](#lambdas)
    - [Branching](#branching)
    - [Arrays](#arrays)
    - [Records](#records)
    - [Strings](#strings)
    - [Console](#console)
    - [Experimental Features](#experimental-features)
- [Usage](#usage)
- [Tests](#tests)
- [Notes on the vision](#notes-on-the-vision)
- [Notes on the design](#notes-on-the-design)
- [Roadmap](#roadmap)
    - [The core language](#the-core-language)
    - [Parse tree tools](#parse-tree-tools)
    - [Nice to haves](#nice-to-haves)
    - [Autonomous program synthesis](#autonomous-program-synthesis)
- [Topics to research](#topics-to-research)


# Language core and syntax

### Expressions

In Gadfly everything is a lexically-scoped _expression_. Gadfly is dynamically
and strongly typed. All values are immutable. An __expression__ is
either a _block_, _lambda_, _predicate_, or _literal_ and all expressions return
a _value_.  __Comments__ begin with the `#` character and continue until the end
of the line. 

A __block__ is a sequence of expressions delimited by a _keyword_ and `end`. A
__keyword__ determines its block's behavior. See the [semantics](#semantics)
section for more details on each keyword. Some examples:

```text
puts "hello world" end

do
  def val "goodbye world" end

  puts val end
end

while rnd < 0.5
  let rnd Math.random end
end

def numbers
  array 1 2 3 end
end

def squares
  map numbers
    fn |n|
      n * n
    end
  end
end

```

### Lambdas, parameters, and arguments

A __lambda__ is a "parameterized block" that is not evaluated until each time it
is called. A lambda can have zero or more _parameters_. A __parameter__ is a
name that is defined each time the lambda is called. Parameters are declared
between `|` characters. If the lambda takes zero parameters, the `|` characters
must be omitted. The  __arguments__ to the lambda are the values of the
expressions in the calling block (using the `@` keyword) bound to the lambda's
parameters. An example:

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
end # => 11

map
  array 1 2 3 end

  fn |n i|
    n + i
  end
end
```

### Predicates, operators, and literals

A __predicate__ is an expression involving an _operator_ and _operands_. See the
[semantics](#semantics) section for more details on each operator. An
__operand__ is either a _predicate_ or a _literal_. A __literal__ is an
expression without subexpressions (string, number, boolean, identifier). A
predicate evaluates to a _number_. `0` is false-y, any other number is truth-y,
and any other value is an error (when used as a boolean). Some examples:

```text
# Not predicates.

fn
  puts "hi" end
end

def val "hi" end

# Predicates.

val

val == "goodbye"

10 > 0

!val
```

_Note that because predicates cannot include blocks they cannot include function
calls. This is somewhat cumbersome to us human programmers, forcing us to write
many instances of trivial indirection, but I think we'll see strong benefits for
code generation and program synthesis. Maybe not, we'll see._

### Variables

A __variable__ is a name that can be resolved to a _value_. A variable is
defined using a `def` block and re-defined using a `let` block. After a variable
is defined it can be referenced in any expression. Some examples

```text
def surname "smith" end

puts surname end

def things
  array
    "thing one"
    "thing two"
  end
end

for things
  fn |thing i|
    puts thing end
  end
end

let things
  push things "thing three" end
end
```

### Values

In `Gadfly` all values are immutable, meaning that all operations on values
return new values and leave the original values unchanged. Note that variables
are not values, they can be redefined (they can be pointed to new values).

The currently supported __value__ types are _string_, _number_, _array_,
_record_,  _lambda_, and _nil_. Strings are delimited by `"` characters. Numbers
are written using decimal notation. The keyword `true` evaluates to `1`, `false`
evaluates to `0`, and `nil` evaluates to `nil`. 

An __array__ is created using the `array` block and  behaves just like a
stereotypical scripting-language array. A __record__ is created using the
`record` block, is a string key to value mapping,  and behaves just like a
stereotypical scripting-language map or dictionary. A __lambda__ is created
using the `fn` block and behaves just like a stereotypical scripting-language
anonymous function that can be passed around and called later.

And that's it for the conventional syntax (e.g. the syntax not relating to
metaprogramming, program synthesis, orchestration, etc.)! The next section
describes the languages keywords and their semantics. After you've read that
you'll be able to write a useful program in `Gadfly`.

# Semantics

In this section we take a look at `Gadfly`'s keywords and their semantics. For
more detailed, runnable examples, see the [examples.core.fly](examples.core.fly)
script. The full set of planned keywords is not yet implemented. _Keywords will
be implemented as needed for the larger goals of the project_.

__In all signatures described below, the `*` character indicates zero or more
occurrences of the preceding expression. The `+` character indicates one or more
occurrences of the preceding expression. The `?` character indicates an optional
expression. Unless otherwise noted, "number", "string", "number", "array", and
"record", and "lambda" are understood to be expressions that evaluate to that
type of value.__

### Variables

`def identifier expression end`

Defines a variable with the given identifier. The variable resolves to the value
of the expression. Variables are _lexically scoped_. If the variable is already
defined in the local scope, it is an error. If the variable is defined in an
outer scope, it will be _shadowed_ in the local scope.

`let identifier expression end`

Re-defines an existing variable with the given identifier. The variable resolves
to the value of the expression. If the variable does not already exist, it is
an error.

### Lambdas

`fn (|identifier+|)? expression end`

When the lambda expression is evaluated, it creates a lambda. They key
difference between a lambda expression and other expressions is that its
subexpressions are not evaluated until the lambda is called. The lambda can take
zero or more parameters. If the lambda takes zero parameters, the `|` characters
must be omitted.

`@ expression* end`

Calls the lambda expression. Each subexpression is evaluated and bound to the
lambda's parameters. The lambda is then evaluated, returning the value of its
last subexpression.

### Branching

The key difference between branching expressions and other expressions is that
their subexpression are evaluated conditionally. The specific behavior of which
expressions are evaluated depends on the keyword.

`if number expression expression end`

If the number is truth-y, the first expression is evaluated. Otherwise, the
second expression is evaluated. The value of the last evaluated expression is
returned.

`and (number expression)+ end`

For each pair of subexpressions, if the first evaluates to a truth-y value, the
second is evaluated. If any of the subexpressions evaluate to a false-y value,
`nil` is returned. Otherwise, the value of the last subexpression is returned.

`or (number expression)+ end`

For each pair of subexpressions, if the first evaluates to a truth-y value, the
second is evaluated and returned. If none of the subexpressions evaluate to a
truth-y value, `nil` is returned.

`while number expression+ end`

While the first expression evaluates to a truth-y value, the rest of the expressions
are evaluated. The value of the last subexpression is returned.

### Arrays

`array expression* end`

Creates an array whose values are the values of the subexpressions. The array is 
returned.

`get array number end`

The value of the array at the index of the number is returned.

`set array number expression end`

Clones the array and sets the value at the index of the number to the value of
the expression. The cloned array is returned.

`for array lambda end`

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
The value of the last evaluated lambda is returned.

`map array lambda end`

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
An array whose values are the result of each lambda call is returned.

`filter array lambda end`

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
An array whose values are the values for which the lambda call returned a
truth-y value is returned.

`reduce array expression lambda end`

For each value in the array, the lambda is called with the value bound to the
lambda's second parameter and the index bound to the lambda's third parameter.
When the lambda is called for the first value in the array, the first parameter
is bound to the value of expression. For each subsequent value in the array, the
first parameter is bound to the value returned by the previous lambda call. The
value of the last evaluated lambda is returned.

`push array expression end`

Clones the array and appends the value of the expression to the cloned array. 
The cloned array is returned.

`pop array end`

Clones the array and removes the last value from the cloned array. The cloned
array is returned.

`unshift array expression end`

Clones the array and prepends the value of the expression to the cloned array.
The cloned array is returned.

`shift array end`

Clones the array and removes the first value from the cloned array. The cloned
array is returned.

`reverse array end`

Clones the array and reverses the order of the values in the cloned array. The
cloned array is returned.

`sort array lambda end`

Clones the array and sorts the values in the cloned array according to the value
returned by the lambda. The lambda takes two parameters, the values of which are
the values in the array. The lambda returns a negative number if the first value
should be sorted before the second, a positive number if the first value should
be sorted after the second, and `0` if the values are equal. The cloned (sorted)
array is returned.

`segment array number number end`

Clones the array and returns a new array whose values are the values of the
cloned array between the first index and the second index (exclusive). The
cloned array is returned.

`splice array number array end`

Clones the first array and divides it in half at the index of the number. It
appends the values of the second array to the first half, and then appends the
second half to the result. The result is returned.

### Records

`record (string expression)* end`

Creates a record whose keys are the strings and whose values are the values of
the expressions. The record is returned.

`read record string end`

The value of the record at the key of the string is returned.

`write record string expression end`

Clones the record and sets the value at the key of the string to the value of
the expression. The cloned record is returned.

`delete record array end`

The array is an array of strings. Clones the record and deletes the keys of the
strings from the cloned record. The cloned record is returned.

`extract record array end`

The array is an array of strings. Returns a record whose keys are the keys of
the strings and whose values are the values of the keys of the strings in the
record. The new record is returned.

`merge record record end`

Clones the first record and then for each kv pair in the second record, sets the
value of the cloned record at the key of the kv pair to the value of the kv
pair. Returns the cloned record.

`keys record end`

An array whose values are the keys of the record is returned.

`values record end`

An array whose values are the values of the record is returned.

### Strings

__TODO__

- [ ] regular expression engine

`split string end`

Returns an array whose values are the characters in the string.

`concat string+ end`

Returns a string whose value is the concatenation of the values of the strings.

`substring string number number end`

Returns a string whose value is the substring of the string between the first
index and the second index (exclusive).

### Console

__TODO__

- [x] `puts`
- [ ] `gets`
- [ ] `err`

`puts expression* end`

Prints the values of the expressions to stdout.

### Experimental Features

_Coming soon!_

- copilot
- distribution
- delegate
- policy
- remote
- memory

And more...


# Usage

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

 - [ ] _WIP_ Happy-path coverage for all keywords and operators
 - [ ] Edge-case coverage for all keywords and operators
 - [ ] Happy-path coverage for at least one robust Gadfly program.
 - [ ] _Down the road_ Fuzzing for all keywords and operators

 Right now we have happy path tests for:

 - [x] array
 - [x] strings
 - [ ] record
 - [ ] branching
 - [ ] lambdas
 - [ ] variables
 - [ ] predicates

 You can run the tests with:

 ```bash
 ./test.sh
 ```

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

# Notes on the design

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

# Roadmap

### The core language

- [x] interpreter
  - [x] lex
  - [x] parse
  - [x] eval
  - [x] closures
- [ ] mvp language features
  - [/] array and associated utilities
  - [ ] hashmaps and associated utilities
  - [/] strings and associated utilities
  - [ ] exceptions and associated utilities
  - [/] puts and gets
  - [ ] http functionality
- [ ] error reporting
- [ ] very simple FFI (maybe)

### Parse tree tools

- [ ] start, stop, pause, resume, retry
- [ ] serialize, deserialize, resume
- [ ] what else?

### Nice to haves

- [ ] syntax highlighter
- [ ] language server protocol implementation
- [ ] repl
- [ ] tail call optimization

### Autonomous program synthesis

- policies
- reflection
- remote subtrees
- generated subtrees

# Topics to research

- probabilistic programming
- compiler design and implementation
- data flow analysis and control flow analysis
- prolog and logic programming
- CSP (communicating sequential processes)
- the actor model
- consensus algorithms
- ...