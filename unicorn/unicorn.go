package unicorn

import (
	"fmt"

	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

type Command struct {
	// Might as well have some code mangling
	// UTF-8 => pre -m flag malformed UTF-8
	Kode bool `help:"Enable kode demangle map mode output (not strict UTF-8)" short:"k"`
	// malformed UTF-8 => formed but mangled instead of strict error marked UTF-8
	Mangle            bool `help:"Enable mangle map mode input (not strict UTF-8)" short:"m"`
	Normal            bool `help:"Enable normalizations (CESU-8 removal, C1 control codes to ^[)" short:"n"`
	UnAscii           bool `help:"Enable ASCII input mapping (to assist upgrading data)" short:"u"`
	clit.StreamFilter      // embedded type .. => .
}

func (c *Command) Help() string {
	return `Unicode mangling depending on the flags.
UTF-8 errors are marked to recover data.`
}

func (c *Command) Run(p *clit.Globals) error {
	// unicorn command hook
	fe.SetGlobals(p)
	fmt.Println(c.InputFile)
	fmt.Println(c.OutputFile)
	return nil
}

// Patterns of unused code points in UTF-8
//
// 0 from 1 byte (7 bit) codes
// ==== ASCII Section Duplicate ====
// 128 from 2 byte (11 bit) codes
// ==== Compact Block ====
// 2048 from 3 byte (16 bit) codes
// ==== Free Plane ====
// 65536 from 4 byte (21 bit) codes
// ==== Free Codes ====
// 15*65536 from 4 byte (21 bit) codes
// ==== High Codes ====
// 32*65536 from 4 byte (21 bit with $F8) codes
//
// ==== PEP 383 ====
// U+DC80...U+DCFF map error conditions
// (U+D800...U+DBFF)...(U+DC00...U+DC7F and U+DD00...U+DFFF)
// (1024-256+128) low only no high prefix
// => as 3 byte
// Given that CESU-8 exists but is discouraged
// and 4 byte encoding of surrogates is "denied" so
// covered by the "Free Plane", this represents
// maximum code point density.
// "Valid" surrogate sequences should be left embedded sometimes?
//
// ==== Control Sequence Introducer ====
// C1 control codes have a duplication via convention of ESC [
// so frees another 32 code points as they were only 3 code points left
// in the Unicode BMP. I think quite a lot of the C0 control codes
// are no longer in common use, as CTRL+C would suggest.
// So that would be an extra 32 in the C1 space
// plus about another 24 in the C0 space.
