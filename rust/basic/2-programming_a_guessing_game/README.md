# Programming a Guessing Game

## Notes

- By default, Rust has a set of items defined in the standard library that it brings into the scope of every program. This set is called the _prelude_. If a type to use isn't in the prelude, it is brought into the scope explicitly with a `use` statement.
- In Rust, variables are immutable by default, meaning once we give the variable a value, the value won't change. To make a variable mutable, we add `mut` before the variable name.
- The `::` syntax before a function, example `String::new`, indicates that the function is an associated function of a type. An _associated function_ is a function that's implemented on a type.
- Like variables, references are immutable by default.
