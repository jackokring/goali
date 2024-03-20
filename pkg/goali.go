//=====================================
//******* Packaging Section ***********
//=====================================

package goali

// PACKAGE DEPENDENCIES
// main	-> goali	-> cli		-> clitype
//					|			-> unicorn	-> filerr	-> clitype
//					|			|			|			-> consts
//					|			|			-> clitype
//					|			-> mickey	-> filerr	-> clitype
//					|			|			|			-> consts
//					|			|			-> clitype
//					|			-> snake	-> filerr	-> clitype
//					|			|			|			-> consts
//					|			|			-> clitype
//					|			-> knap		-> filerr	-> clitype
//					|						|			-> consts
//					|						-> clitype
//					-> filerr	-> clitype
//					|			-> consts
//					-> consts
//					-> gin		-> filerr	-> clitype
//											-> consts

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"

	"log"
	"log/syslog"

	"github.com/jackokring/goali/cli"
	con "github.com/jackokring/goali/consts"
	fe "github.com/jackokring/goali/filerr"
	"github.com/jackokring/goali/gin"
)

//=====================================
//********** Main Section *************
//=====================================

// Find the system configuration directory
func SystemConfigDir() string { // Linux
	// yes you're crazy configuration sets are a NO!
	// i guess the first is the best as per $PATH
	// and it's not an over merge apply of last to first
	systemConfig := strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")
	if len(systemConfig) == 0 { // Windows
		systemConfig = []string{os.Getenv("PROGRAMDATA")}
		if len(systemConfig[0]) == 0 { // MacOS
			systemConfig[0] = "/Library/Application Support"
		}
	}
	// should be fine on a well configured system
	return systemConfig[0]
}

// # Main Entry Point
func Goali() {
	// full config loading
	// the pro-file sub tree can be supplied from a file on the CLI
	//globalConfig := "/etc/" + AppName + "/config.yaml"
	dir, err := os.UserConfigDir() // {
	// should the error handler go here syntax wise??
	// tuple implicit?
	// }
	if err != nil { // Error(err) not used as not critical
		dir2, err2 := os.UserHomeDir()
		// pretty critical to have a home directory?
		// maybe some sort of demon process?
		// Fatal(err2)
		dir = dir2
		if err2 != nil {
			// pretends to be Darwin on failing
			dir = SystemConfigDir()
		}
	}
	localConfig := filepath.Join(dir, "."+con.AppName+".yaml")
	// Now we can parse
	ctx := kong.Parse(&cli.Cli,
		kong.Configuration(kongyaml.Loader /* globalConfig, */, localConfig),
		kong.Vars{
			// ${<name>} in `tags`
			//"globalConfig": globalConfig,
			"localConfig": localConfig,
			"appName":     con.AppName,
			"maxVerbose":  strconv.Itoa(fe.MaxVerbose),
		},
		// loading defaults for flags and options
		kong.NamedMapper("yamlfile", kongyaml.YAMLFileMapper),
		kong.Description(con.AppDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}),
	)
	log.SetOutput(os.Stderr)
	debug := 0
	if cli.Cli.Debug {
		debug += log.Lshortfile | log.Lmicroseconds
	}
	log.SetFlags(log.LstdFlags | log.LUTC | debug)
	//Error(errors.New("Error test"))
	if cli.Cli.SysLog {
		// Configure logger to write to the syslog.
		logwriter, e := syslog.New(syslog.LOG_NOTICE, con.AppName)
		fe.Fatal(e)
		log.SetOutput(logwriter)
	}
	fe.Lock.Lock() // lock the IO temporarily
	gin.Tui()

	// Call the Run() method of the selected parsed command.
	// Extra context arg as not cast to command
	fe.Fatal(ctx.Run(&cli.Cli.Globals))
	// So you've found an Error?
	// Have you considered using the functions:
	//
	// 		func Error(err) bool // in an if test handler
	//		func Fatal(err) // anywhere
	//		func Hard(err)	bool // anywhere
	//
	// As these will provide panic info with the -d option.
	//
	//		func (*command)Run(*clitype.Globals) error
	//
	// Yes, returning a nil is an option. Source code error?

	// Are the messages on a separate thread?
	// Tea might channel lock if ctx.Run() did not
	// cause fe.Lock.Unlock() to be called.
	// If the IO was not unlocked (by accident)
	// then the tea program is not running.
	// If the IO was unlocked the tea program has "likely" started.
	// In a rare case it will be waiting on the fe.Lock ...
	time.Sleep(250 * time.Millisecond) // give a waiter a try?
	if !fe.Lock.TryLock() {            // can not grant so shut down error
		fe.Fatal(fmt.Errorf("the file lock needs unlocking in the command run function"))
	}
	gin.Tea(gin.QuitMsg{})
	var finalModel gin.Model
	var okToExit bool = false
	for !okToExit {
		finalModel, okToExit = gin.TuiGetModel()
	}
	ra := finalModel.RunAfter
	if ra != nil {
		ra.RunAfter() // post TUI model postAction receiver
	}
	fe.CloseAll(false) // natural exit
}
