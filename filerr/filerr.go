package filerr

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"

	clit "github.com/jackokring/goali/clitype"
	con "github.com/jackokring/goali/consts"
)

//=====================================
//****** Error Handler Section ********
//=====================================

// The exit code type
type ExitCode int // uint8

// The primitive exit codes
const (
	// general error (used by kong errors, and other library errors)
	ERR_GENERAL ExitCode = 1 << iota
	// all fatal errors have this set (all non-fatal errors become this with -x option)
	ERR_FATAL
	// 4 bits of error code with flags in lover mask
	ERR_0
	//
	ERR_1
	//
	ERR_2
	//
	ERR_3
	// for some shell errors lower than 128 (64+n)
	ERR_SHELL
	// internal code (128+n) for signals like SIGHUP, SIGTERM etc.
	ERR_SIGNAL_HANDLER
	// (> uint8) maybe used but, some applications and the shell might not support the code
	ERR_RANGE_PLUS_ONE
)

const ( // maximum of 16 possible bit patterns before "shell" overflow
	// basic error code space for more specific errors
	ERR_RESET_UNCLASSIFIED ExitCode = iota << 2
	ERR_STREAM
	E_02
	E_03
	// ...
	E_10
	E_11
	E_12
	E_13
	//
	E_20
	E_21
	E_22
	E_23
	//
	E_30
	E_31
	E_32
	E_33
)

// The combination exit codes useful named list
const (
	ERR_MINUS_ONE     ExitCode = ExitCode(^0)                   // two's complement inversion
	ERR_SIGNAL_CTRL_C          = ERR_SIGNAL_HANDLER | ERR_FATAL // for example
	ERR_RANGE                  = ERR_RANGE_PLUS_ONE - 1         // 255 (technically also ERR_MINUS_ONE)
	ERR_WRONG                  = ERR_FATAL | ERR_GENERAL        // both as general non fatal made fatal
	ERR_SHELL_CMPLX            = ERR_SIGNAL_HANDLER | ERR_SHELL // both for some complex shell errors (196+n)
)

// A concrete extended error type
type E struct {
	error
	exit ExitCode
}

// The extended error cast interface
type R interface {
	error
	Exit()
	With(exit ExitCode) R
}

// The simplest contract number
func (e E) Exit() {
	os.Exit(int(e.exit)) // bumped up to typed int for 32 set entries
}

// Set the exit code
func (e *E) With(exit ExitCode) R {
	if exit == ERR_RESET_UNCLASSIFIED { // allows extended codes by lower bits reset
		e.exit = e.exit & ^3 // clear lower bits
	}
	e.exit = exit | e.exit
	return e
}

var g *clit.Globals

// Called first by any Command to allow access of the global switches.
func SetGlobals(globals *clit.Globals) {
	g = globals
}

// notify the current logger writer.
//
// Use a nest of 0 for the details of the caller of notify()
func notify(s string, nest int) {
	if g.Quiet {
		return
	}
	if g.SysLog || os.Stderr != os.Stdout { // external so OK as no collide TUI
		// - STDOUT target situation barrier
		// auto quiet
		// caller of notify() skips 2 stack frames
		log.Output(nest+2, s)
	}
}

// Error not nil syntax sugar.
func Error(e error) bool {
	// better naming but should var args limit a callback?
	// indicated in a function signature by
	// func (rx type) name (args) { body } {}
	// you know, as an empty {} ... filled on use?
	//
	if e != nil {
		// and a level of nest for caller of Error()
		if g.Wrong {
			fatalNest(&E{e, ERR_WRONG}, 1) // caller of Error
			return false                   // never happens
		}
		// print once
		notify(e.Error(), 1) // {} here handler
		return true
	}
	return false
}

// The array of things to close
var closers []io.Closer = []io.Closer{}

// CloseAll open things.
func CloseAll(roll bool) {
	// apparently logical for closing the outer streams first
	for i, j := 0, len(closers)-1; i < j; i, j = i+1, j-1 {
		closers[i], closers[j] = closers[j], closers[i]
	}
	for _, c := range closers {
		w, ok := c.(FilterWriter)
		if ok {
			w.Rollback(nil)
		}
		Error(c.Close())
	}
}

// DeferClose of an open thing.
func DeferClose(c io.Closer) {
	closers = append(closers, c)
}

// Fatal error logging.
func Fatal(e error, x ...ExitCode) {
	if e == nil {
		return
	}
	c := &E{
		e, ERR_FATAL, // essential to have non-zero default
	}
	for _, ec := range x {
		c.With(ec)
	}
	if c.exit == ERR_RESET_UNCLASSIFIED { // check success by accident
		c.With(ERR_FATAL) // make fatal again
	}
	fatalNest(c, 1) // one stack frame for the proxy call
}

// FatalNest with skipped stack frames
func fatalNest(e R, skip int) {
	//goali.Tea(gin.QuitMsg)
	// As a panic or exit due anyway
	// the close is to allow file write
	// sync to already done options
	// No notify() proxy as serious terminal error
	CloseAll(g.Rollback) // rollback
	if g.Debug {
		// this should always drop somewhere
		log.Panic(e.Error()) // your basic panic, with no error code options
	}
	// stack skip Output and Fatal (2)
	log.Output(skip+2, e.Error())
	e.Exit()
	// a botch general error
	// technically this should be a uint8, but these days of 32 bit adults ...
	// and 64 bit kids ...
}

// Notify a debug message to the current logger writer.
func Debug(s string) {
	if g.Debug {
		notify(s, 1) // the caller of Debug()
	}
}

//=====================================
//**** File Abstraction Section *******
//=====================================

// GetIO
//
// Open reader before writer for error
// interaction effect with STDOUT by "-".
// This order make sense in the context
// of a "future" optimizing compiler
// where the read file is reusable immediately
// and the write needs to be closed to commit
// as maybe it would support "rollback"
// on Rollback() with the replacement
// happening atomically on the Close().
func GetIO(io clit.StreamFilter) (FilterReader, FilterWriter) {
	return GetReader(io.InputFile), GetWriter(io.OutputFile)
}

// Sure I need an GetIORW(io IoFile) (FilterReader, FilterWriter).
//
// It might just come in useful. Be aware that the output is a zero length file to which
// you can io.Copy to if you do need an exact clone to start. This may cause problems
// as not all file kinds can then Seek back to the beginning.
func GetRW(io clit.IoFile) (FilterReader, FilterWriter) {
	w := GetWriter(clit.OutputFile{
		Compress:   io.Compand,
		Force:      true,
		Group:      io.Group,
		OutputFile: io.IoFile,
		Write:      io.Write,
	})
	// use backup as input file
	n := w.getRollback()
	if n == "" { // there is no rollback file
		Fatal(fmt.Errorf("can't construct input file from old output file content"))
	}
	// see Rollback(closeBefore FilterReader)
	r := GetReader(clit.InputFile{InputFile: n, Expand: io.Compand}) // closed first
	return r, w
}

// FilterReader is an abstraction to allow the wrapped
// unfiltered streams to be closed possibly by cascade calling.
type FilterReader interface {
	io.Closer
	// io.EOF
	Read(b []byte) int
	EOF() bool
	io.Seeker
}

// A concrete GZip FilterReader
type GReader struct {
	this io.ReadCloser
	// is it a Closer => this == nil
	// the wrapped or inner Closer
	wrap io.ReadCloser
	// requires pointer receiver
	// so all instances must be
	// by address &.
	thisEof bool
}

func (r GReader) Close() error {
	Error(r.this.Close())
	if r.wrap != nil {
		Error(r.wrap.Close())
	}
	// already handled display of errors
	return nil
}

// N.B. Due to needing to alter the EOF state
// the *GReader becomes required. This causes
// a cascade to need all interface instances
// to need pointer to value by &. Otherwise
// how would the pointer refer to that
// which is to be modified?
func (r *GReader) Read(b []byte) int {
	n, e := r.this.Read(b)
	if e == io.EOF {
		// delay spec for while style test of EOF
		r.thisEof = true
	} else {
		Fatal(e)
	}
	return n
}

func (r GReader) EOF() bool {
	return r.thisEof
}

func (r GReader) Seek(offset int64, whence int) (int64, error) {
	s, ok := r.this.(io.Seeker)
	if ok {
		n, e := s.Seek(offset, whence)
		return n, e
	}
	return 0, fmt.Errorf("reader seek not allowed")
}

// FilterWriter is an abstraction to allow the wrapped
// unfiltered streams to be closed possibly by cascade calling.
type FilterWriter interface {
	io.Closer
	// io.EOF? on writing?
	Write(b []byte) int
	// useful for RW paradigm
	getRollback() string
	// rollback future
	Rollback(closeBefore FilterReader)
	io.Seeker
}

// A concrete GZip FilterWriter
type GWriter struct {
	this io.WriteCloser
	// is it a Closer => this == nil
	// the wrapped or inner Closer
	wrap io.WriteCloser
	// rollback temp filename
	rollback string
	mode     fs.FileMode
	out      string
}

func (w GWriter) Close() error {
	Error(w.this.Close())
	if w.wrap != nil {
		Error(w.wrap.Close())
	}
	if w.rollback != "" {
		// remove roll back up
		Error(os.Remove(w.rollback))
	}
	// already handled display of errors
	return nil
}

func (w GWriter) Write(b []byte) int {
	i, e := w.this.Write(b)
	Fatal(e)
	return i
}

// Rollback the writer and allow closing an associated reader (this can be null).
func (w GWriter) Rollback(closeBefore FilterReader) {
	// as interface (nil, nil)
	if closeBefore != nil {
		Error(closeBefore.Close())
	}
	Error(w.this.Close())
	if w.wrap != nil {
		Error(w.wrap.Close())
	}
	if w.rollback != "" {
		flags := os.O_WRONLY | os.O_CREATE | os.O_EXCL
		Fatal(os.Remove(w.out))
		in, e := os.Open(w.rollback)
		Fatal(e)
		out, e2 := os.OpenFile(w.out, flags, w.mode)
		if e2 != nil {
			Error(in.Close())
			Fatal(e2)
		}
		_, e3 := io.Copy(out, in)
		if e3 != nil {
			Error(in.Close())
			Error(out.Close())
			Fatal(e3)
		}
		Fatal(os.Remove(w.rollback))
	}
}

func (w GWriter) getRollback() string {
	// remove cast exception
	return w.rollback
}

func (w GWriter) Seek(offset int64, whence int) (int64, error) {
	s, ok := w.this.(io.Seeker)
	if ok {
		n, e := s.Seek(offset, whence)
		return n, e
	}
	return 0, fmt.Errorf("writer seek not allowed")
}

// Get reader
func GetReader(i clit.InputFile) FilterReader {
	r := getReaderBase(i)
	r.Read(make([]byte, 0)) // sure needs a zero length EOF set
	// I prefer a `for !EOF {}` kind of thing
	// ... do {} while(); ... C-Form???????
	return r
}

func getReaderBase(i clit.InputFile) FilterReader {
	if i.InputFile == "-" {
		in := os.Stdin
		nin, e := os.Open(os.DevNull)
		Fatal(e)
		os.Stdin = nin
		DeferClose(in)
		return &GReader{in, nil, false}
	}
	f, err := os.Open(i.InputFile)
	Fatal(err)
	DeferClose(f)
	if i.Expand {
		f2, err2 := gzip.NewReader(f)
		Fatal(err2)
		DeferClose(f2)
		return &GReader{f2, f, false}
	}
	return &GReader{f, nil, false}
}

// Get writer
func GetWriter(o clit.OutputFile) FilterWriter {
	if o.OutputFile == "-" {
		out := os.Stdout
		// Handle TUI expectations
		os.Stdout = os.Stderr
		// already -q as command may have notify()
		// on logger mixing
		DeferClose(out) // just in case pipe
		return GWriter{out, nil, "", 0, ""}
	}
	// create if not exist <- N.B.
	var perms fs.FileMode = 0644
	if o.Write && !o.Group {
		perms = 0664 // give group permissive permissions
	}
	if o.Group && !o.Write {
		perms = 0640 // remove everybody permissions
	}
	if o.Group && o.Write {
		perms = 0660
	}
	flags := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	var rollback string
	m, ex := os.Stat(o.OutputFile)
	mode := m.Mode()
	if o.Force && ex == nil { // and exists, else no backup
		// of course the "future" compiler would
		// have to insist on supplying a force
		// "open" token here, for a possible
		// commit vs. rollback.
		// make backup?
		r, e := os.Open(o.OutputFile)
		Fatal(e)
		w, e2 := os.CreateTemp("", con.AppName+"-*.bak")
		if e2 != nil {
			Error(r.Close())
			Fatal(e2)
		}
		_, e3 := io.Copy(w, r)
		if e3 != nil {
			Error(r.Close())
			Error(w.Close())
			Fatal(e3)
		}
		Error(r.Close())
		Error(w.Close())
		Fatal(os.Remove(o.OutputFile))
		rollback = w.Name()
		// Backed up!
	}
	f, err := os.OpenFile(o.OutputFile, flags, perms)
	Fatal(err)
	DeferClose(f)
	if o.Compress {
		f2, err2 := gzip.NewWriterLevel(f, gzip.BestCompression)
		Fatal(err2)
		DeferClose(f2)
		return GWriter{f2, f, rollback, mode, f.Name()}
	}
	return GWriter{f, nil, rollback, mode, f.Name()}
}
