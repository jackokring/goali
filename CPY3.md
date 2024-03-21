# CPY3 is a submodule

`git submodule update --init`

## Current Version Python 3.11 (2024-03-21)

The main `go.mod` file has a `replace` for the submodule. Only critical references
to the repository name have been changed. So no, the examples have not been changed.

## Generic `pkg-config python3-embed`

This will help in the future.

## `Py_AddCFunction` Singular

This is a much better idea for CLI functionality injection via flags.
A nice addition essential to "field" the standard IO. You may ask yourself why not just do a python
module with C, or some go exporting C with a python wrapper?

Well, `python` the binary is but a C program, and in a strange sort of `flex`/`bison` way `go`
is a compiled translation client into C too.