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
