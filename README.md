# gadfly

__gadfly__ is an experimental programming language and treewalk interpreter
designed for autonomous program synthesis. To this end, the core language is
intended to be very simple, very regular, and amenable to certain kinds of
analysis and metaprogramming. It is heavily inspired by
[scheme](https://www.scheme.org/) and [ruby](https://www.ruby-lang.org/en/). It
is dynamically and strongly typed. 

_Please note, this project and documentation are under heavy development. If you
see something is missing, find an error, have a question, or have anything at
all to say, an [issue](https://github.com/killthebuddh4/gadfl.ai/issues) would
be awesome!_

# contents

- [gadfly](#gadfly)
- [contents](#contents)
- [language core and syntax](#language-core-and-syntax)
    - [expressions](#expressions)
    - [lambdas, parameters, and arguments](#lambdas-parameters-and-arguments)
    - [predicates, operators, and literals](#predicates-operators-and-literals)
    - [variables](#variables)
    - [values](#values)
- [semantics](#semantics)
    - [variables](#variables-1)
    - [lambdas](#lambdas)
    - [branching](#branching)
    - [loops](#loops)
    - [logging](#logging)
    - [arrays](#arrays)
    - [records](#records)
    - [strings](#strings)
    - [experimental](#experimental)
- [notes on the vision](#notes-on-the-vision)
- [notes on the design](#notes-on-the-design)
- [roadmap](#roadmap)
    - [the core language](#the-core-language)
    - [parse tree tools](#parse-tree-tools)
    - [nice to haves](#nice-to-haves)
    - [autonomous program synthesis](#autonomous-program-synthesis)
- [topics to research](#topics-to-research)


# language core and syntax

### expressions

In gadfly everything is a lexically-scoped _expression_.  An __expression__ is
either a _block_, _lambda_, _predicate_, or _literal_ and all expressions yield
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

### lambdas, parameters, and arguments

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

### predicates, operators, and literals

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

### variables

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

### values

In `gadfly` all values are immutable, meaning that all operations on values
return new values and leave the original values unchanged. Note that variables
are not values, they can be redefined (they can be pointed to new values).

The currently supported __value__ types are _string_, _float_, _array_,
_record_,  _lambda_, and _nil_. Strings are delimited by `"` characters. Floats
are written using decimal notation. The keyword `true` evaluates to `0`, `false`
evaluates to `1`, and `nil` evaluates to `nil`.

An __array__ is created using the `array` block and  behaves just like a
stereotypical scripting-language array. A __record__ is created using the
`record` block, is a string key to value mapping,  and behaves just like a
stereotypical scripting-language map or dictionary. A __lambda__ is created
using the `fn` block and behaves just like a stereotypical scripting-language
anonymous function that can be passed around and called later.

And that's it for the conventional syntax (e.g. the syntax not relating to
metaprogramming, program synthesis, orchestration, etc.)! The next section
describes the languages keywords and their semantics. After you've read that
you'll be able to write a useful program in `gadfly`.

# semantics

In this section we take a look at `gadfly`'s keywords and their semantics. For
more detailed, runnable examples, see the [examples.core.fly](examples.core.fly)
script. The full set of planned keywords is not yet implemented. _Keywords will
be implemented as needed for the larger goals of the project_.

### variables

### lambdas

### branching

### loops

### logging

### arrays

### records

### strings

### experimental

_Coming soon!_

- copilot
- distribution
- delegate
- policy
- remote
- memory

And more...


_Requires `go` 1.21 or higher. Learn how to install `go` [here](https://go.dev/doc/install)._

```bash
go run . <path to gadfly source>
```

```bash
go run example.fizzbuzz.gadfly
go run example.sieve.gadfly
go run example.fibonacci.gadfly
go run example.factorial.gadfly
go run example.palindrome.gadfly
```

# notes on the vision

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

# notes on the design

- It's looking like the language will (unsurprisingly) be very Lispy. One way to
  think about things is that `gadfly` takes homoiconicity to wild extremes.
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

# roadmap

### the core language

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

### parse tree tools

- [ ] start, stop, pause, resume, retry
- [ ] serialize, deserialize, resume
- [ ] what else?

### nice to haves

- [ ] syntax highlighter
- [ ] language server protocol implementation
- [ ] repl
- [ ] tail call optimization

### autonomous program synthesis

- policies
- reflection
- remote subtrees
- generated subtrees

# topics to research

- probabilistic programming
- compiler design and implementation
- data flow analysis and control flow analysis
- prolog and logic programming
- CSP (communicating sequential processes)
- the actor model
- consensus algorithms
- ...