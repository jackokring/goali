package knap

//=====================================
//********* Knap Web Servia ***********
//=====================================

import (
	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

type Command struct {
}

func (c *Command) Help() string {
	return `Web Servia.`
}

func (c *Command) Run(p *clit.Globals) error {
	// knap command hook
	fe.SetGlobals(p)

	return nil
}
