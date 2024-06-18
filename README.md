# Goali

Not to be confused with the non-existent marsh-soup-eal the Koali. :D
My `init.el` file for [Emacs][emacs] is growing to replace VSCode.

## Install
```
$ # you could use godeb instead of default golang
$ sudo apt install postgres
$ sudo apt install git python-is-python3 golang python3-pip python3-dev libxxf86vm-dev
$ git clone git@github.com:jackokring/goali.git
$ cd goali
$ # create and restore db
$ ./restore.sh
$ # you may find the sudo for package dependencies in require.sh in a comment
$ # the cpy3 submodule for embedding python
$ # pull and satisfy requirements then go build
$ ./require.sh
```

## Fork
```
$ # go to repository directory
$ cd goali
$ # use your fork URL
$ git remote set-url origin git@github.com:<username>/goali.git
$ git push origin master
$ # perhaps dump the database
$ ./dump.sh
$ # check requirements and git add commit push
$ ./freeze.sh
```
And then create pull requests if you feel like it.

## Links

* [Goali Blog][blog] A github pages site. Keeps the speculation off this readme.
* [TODO bot][todo] A github bot using commented `@todo`/`@body` to raise automatic issues.
* [A Goali][David] This repository has nothing to do with footie or crisps.

## `goali` Main Commands

* `unicorn` Unicode mangler. Also [The Word][the_word] and what were if [6602][6602]...
* `snake` Embedded and expanded python...
* `mickey` Nice GUI for something...
* `knap` Web servia thing...

**Under Construction TODO...**

## Features

* `--pro-file=PROFILE` process configuration override with `.yaml` files.
* Logging and `os.Stderr` fallback.
* `os.Stdin`/`os.Stdout` use by `-` filename.
* GZip IO wrapping CLI options.
* Use `c`, `python` and/or `go` to develop possibilities.

## Coding Framework Features

* Debug 'n panic `Fatal(err) void`/`Error(err) bool` framework.
* `FilterWriter.Rollback()` for a non-commit `Close()`.

---

## Why Go?

So 26 MB instead of a rust 200 MB. Cool for a demo of first build.
Much faster code compile too.

`go build -ldflags "-w -s"` saves 5 MB using the linker optimizer.

## Go Modules Used (and Their Indirect Modules)

* [kong][kong] - CLI parse
* [fyne][fyne] - GUI toolkit
* [chi][chi] - A kind of node.js express router in go
* [cpy3][cpy3] - Python 3.11 wrapper `git submodule update --init --recursive` ([CPY3.md][CPY3])
* [bubbletea][bubbletea] - TUI toolkit
* [lipgloss][lipgloss] - ANSI sequences
* [bubbles][bubbles] - TUI extended controls
* [kong-yaml][kong-yaml] - YAML config loader
* [squirrel][squirrel] - SQL statement builder
* [pgx][pgx] - Postgres connector
* [go-keyring][go-keyring] - Secret keyring handler (needs local `login` keyring via `seahorse`)
* [godeb][godeb] - A go version automatic `.deb` maker (`godeb`) to install go versions (`go install gopkg.in/niemeyer/godeb.v1/cmd/godeb`)
* [expr][expr] - An expression language
* [govalidator][govalidator] A data format validator
* [sqlc][sqlc] An SQL automatic code generation wrapper 

## Modules to Find

This place is kind of a brain storming section, stuff in the decision matrix.

* [Awesome Go with Stars][Awesome]

## Python `venv`

A python virtual environment was added to the project. Its major directories
`lib`, and `bin` were added to `.gitignore`. This may be altered later 
depending on functional use. A `snake_test.ipynb` test Jupyter notebook is included.

---

## Go Likes and Dislikes

### Likes

* It's fast enough at compile and doesn't clog my Chromebook.
* Receivers.
* No `switch` break madness fall through.
* Implicit interface implementations, without declaration and locking out
inbuilt type extension.
* Allows submodules and redirection via `replace ... => ...` to local a `git submodule`.
* Initial capitalization exports.
* `.(type)` switch variants. I like the two assignment tuple form.
* A mindset of default reentrant non-atomicity.
* Embedded `struct` types.
* Ah, `chan` of Occam. Or should that've been Ockham or Oakham? Nice `select` for
`ALT` too.
* The `error` strategy. Sure I've hacked it (see `filerr/filerr.go`).
* Although the idea of tacit tuple receivers excites me, they're not needed. 

### Dislikes

* The postfix typing. As from a point of view it has no parse introduction of
type information **before** an identifier. This slightly increases parse complexity
and prevents some kind of edit selection of type DropDown inserting an
identifier TextBox. It's not a deal breaker, as I've seen / used Pascal and
Modula-2 / Oberon before.
* The weird `import "github.com/jackokring/v2/cpy3"` and making a new `v2` branch
along with a `module github.com/jackokring/cpy3/v2` just to allow a
`go get github.com/jackokring/cpy3/v2@v2.0.0` after a following
`GOPROXY=proxy.golang.org go list -m github.com/jackokring/cpy3/v2@v2.0.0` but
only after a `git tag v2.0.0` itself after a push followed by a `git push origin v2.0.0`.
* Of course the above is made more irritating by node.js inside VSCode spamming
the `-jMAX` option and behaving in its docs like the other processes on the
system are the problem. Apparently, the terminal failure to initialize on first go
is a feature, and it can only count to 120. It doesn't understand `tmux`?
This is kind of "fixed" using VSCode as the terminal (apart from SIGHUP on close).
At least it has split views.
* The segmentation fault on `errors.New()`. But, it was new. Likely a pointer
to local returned from function and dereferenced.
* A mild `.cache` directory half a GB of "junk" from the language server. Perhaps
a `go tidy` for a chop down project global.
* The `import` deletion from a `^S` when there is an intent from `go get` and
no `go build` has been issued. I mean yes it does suggest and add, but does
it get the preferred named as data back (project or directory wide)?
* Mild annoyance with anonymous function closure syntax for `return` to achieve a
multi `break` and not even an ORIC-1 BASIC `POP:RETURN`. I'm not a fan of the
named labels and `goto` approach. Perhaps a `func {}` short form without `()`.
I mean I can't suggest a `return[level]` syntax as `[]int{x, y, z ...}` might
mix bad with the parser, being a prefixed typing.

### Wonderings

* Is `const` (as a type read only intent) inferred for auto `VAR` arguments (Pascal)?

### Things which are Really VSCode Things

* Eager launching of resource using tools. I'd like an awareness of `-j` which
does not use all my cores, as you know it's just a Debian container
also running a browser, with the music, and the reference internet. It
should hold off when close to memory 70% or core usage above 50% like
give the syntax check a stop for a minute or two. I only deleted a `}`. I'm sure
I'll manage without the extra or fewer items in the "intelli-drop-list".
* Aren't I lucky I didn't install the C tools recommended `-jINFINTE` balls.
Sure the `mc` in da house.
* `pylance` maybe a slightly bigger boil than it need be.
* Yes, the environment slows down when altering ANY cgo file.
* `VSCodeLIght` with less gfx buffers? Oh, that fancy fast compositor
gigabyte depositor. But sorry, "BOINC" is more appropriate. You know
"useful" sh*t @ home?
* OK, as a prototype it kind of works ;D. Tai-Ping? Is he 'ard lady boy?

### Things which are Really Sinter Net Things

* Ah, hot metal. Maybe a proxy JS wall and `*.md` server? Like for local
`*.html` building? Some user control data sharding? It obviously needs some
authority based access protocols for utility. `_O_-_ETA_` yeah, like, bean
ear lawn?

---

## Extracted Link Definitions (Invisible)

[blog]: jackokring.github.io/goali
[bubbles]: github.com/charmbracelet/bubbles
[bubbletea]: github.com/charmbracelet/bubbletea
[chi]: github.com/go-chi/chi/v5
[cpy3]: github.com/jackokring/cpy3
[expr]: github.com/expr-lang/expr
[go-keyring]: github.com/zalando/go-keyring
[godeb]: gopkg.in/niemeyer/godeb.v1/cmd/godeb
[govalidator]: github.com/asaskevich/govalidator
[fyne]: fyne.io/fyne/v2
[kong]: github.com/alecthomas/kong
[kong-yaml]: github.com/alecthomas/kong-yaml
[lipgloss]: github.com/charmbracelet/lipgloss
[pgx]: github.com/jackc/pgx
[sqlc]: github.com/sqlc-dev/sqlc
[squirrel]: github.com/Masterminds/squirrel
[todo]: todo.jasonet.co

[Awesome]: github.com/amanbolat/awesome-go-with-stars
[CPY3]: CPY3.md
[David]: en.wikipedia.org/wiki/David_Seaman

[the_word]: docs.google.com/document/d/1rsPyq3c7uVzxpUb9JXtq0b603HSuju7NWeZ_aYfVkzs/edit?usp=sharing
[6602]: docs.google.com/spreadsheets/d/1ejnimh5PPYHhPX-k93yTYFoXf0_bTxylGQEbKL1TbAw/edit?usp=sharing
[emacs]: github.com/jackokring/goali/blob/master/extras-backup/init.el