# gadfly

`gadfly` is an experimental programming language and treewalk interpreter,
written in `go`, and designed for autonomous program synthesis. _It is currently
under heavy development._

# notes on the vision

- Imagine a programming language designed from the ground up to be used by
  language models (LMs).
- An LM wouldn't need to generalize, it could write a new program for each new
  task it encounters.
- An LM user could analyze running programs very quickly, it could modify
  running programs according to new data.
- Instead of importing code, an LM user could search for similar code in a
 . database and then repurpose it.
- Because an LM user could analyze running programs very quickly, it could make
  sense to store all kinds of useful metadata in the parse tree.
- More generally, the parse tree could be something that is constantly
  manipulated by the LM user.
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

# basic usage

# basic examples

# roadmap

### the core language

- [ ] interpreter
  - [ ] lex
  - [ ] parse
  - [ ] eval
- [ ] mvp language features
  - [ ] array and associated utilities
  - [ ] hashmaps and associated utilities
  - [ ] strings and associated utilities
  - [ ] http functionality
- [ ] error reporting
- [ ] very simple FFI (maybe)

### parse tree tools

- [ ] start, stop, pause, resume, retry
- [ ] serialize, deserialize, resume
- [ ] what else?

### autonomous program synthesis

...






