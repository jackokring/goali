package cli

import (
	clit "github.com/jackokring/goali/clitype"
	"github.com/jackokring/goali/knap"
	"github.com/jackokring/goali/mickey"
	"github.com/jackokring/goali/snake"
	"github.com/jackokring/goali/unicorn"
)

//=====================================
//****** CLI Structure Section ********
//=====================================

// optional:""
// type:"path"	A path. ~ expansion is applied. - is accepted for stdout, and will be passed unaltered.
// type:"existingfile"	An existing file. ~ expansion is applied. - is accepted for stdin, and will be passed unaltered.
// type:"existingdir"	An existing directory. ~ expansion is applied.
// type:"counter"	Increment a numeric field. Useful for -vvv. Can accept -s, --long or --long=N.
// type:"filecontent"	Read the file at path into the field. ~ expansion is applied. - is accepted for stdin, and will be passed unaltered.
// env:"X,Y,..."	Specify envars to use for default value. The envs are resolved in the declared order. The first value found is used.
// default:"X"	Default value.
// short:"X"	Short name, if flag.
// negatable:""	If present on a bool field, supports prefixing a flag with --no- to invert the default value.
// envprefix:"X"	Envar prefix for all sub-flags.
// passthrough:""	If present on a positional argument, it stops flag parsing when encountered, as if -- was processed before. Useful for external command wrappers, like exec. On a command it requires that the command contains only one argument of type []string which is then filled with everything following the command, unparsed.

// Flag allocation matrix fio and goa add to any tool with file IO and global options
//flg  goa fio uni mic
// a
// b
// c		x	x
// d	x		x	x
// e		x	x
// f		x	x	x
// g		x	x
// h	x 		x	x
// i
// j
// k			x
// >> L was almost going to be monotone ^G jingle bells.
// l
// m			x
// n			x
// o
// p	x		x	x
// q	x		x	x
// r		x
// s	x		x	x
// t	x
// u			x
// v	x		x	x
// w		x	x
// x	x		x	x
// y
// z

// Alphabetically sorted command line arrangement.
var Cli struct {
	clit.Globals
	// see individual named Command packages
	Knap    knap.Command    `cmd:"" help:"Web servia."`
	Mickey  mickey.Command  `cmd:"" help:"GUI launcher."`
	Snake   snake.Command   `cmd:"" help:"Python interpreter."`
	Unicorn unicorn.Command `cmd:"" help:"Unicode mangler."`
}
