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

// The 100% verbose level.
const MaxVerbose = 3

var g *clit.Globals

// Called first by any Command to allow access of the global switches.
func SetGlobals(globals *clit.Globals) {
	g = globals
}

// Notify the current logger writer.
func Notify(s any) {
	if g.Quiet {
		return
	}
	if g.SysLog { // external so OK as no collide TUI
		log.Print(s)
	} else { // external not used
		// - STDOUT target situation barrier
		// auto quiet
		if os.Stderr != os.Stdout {
			log.Print(s)
		}
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
		Notify(e) // {} here handler
		return true
	}
	return false
}

// The array of things to close
var closers []io.Closer = []io.Closer{}

// CloseAll open things.
func CloseAll() {
	// apparently logical for closing the outer streams first
	for i, j := 0, len(closers)-1; i < j; i, j = i+1, j-1 {
		closers[i], closers[j] = closers[j], closers[i]
	}
	for _, c := range closers {
		Error(c.Close())
	}
}

// DeferClose of an open thing.
func DeferClose(c io.Closer) {
	closers = append(closers, c)
}

// Fatal error logging.
func Fatal(e error) {
	if Error(e) {
		if g.Debug {
			// this should always drop somewhere
			log.Panic(e)
		}
		//
		// No Notify() proxy as serious terminal error
		CloseAll()
		log.Fatal("FATAL: ", e)
	}
}

// Hard error check logging.
//
// Check to see if the hard error flag is set
// and cause all errors to stop the application.
func Hard(e error) bool {
	if g.Wrong {
		Fatal(e)
		return false
	}
	return Error(e)
}

// Notify a debug message to the current logger writer.
func Debug(s any) {
	if g.Debug {
		Notify(s)
	}
}

// Verbose measure of the logged output status to show.
func Verbose() int {
	if g.Quiet { // quiet or STDOUT priority?
		return 0
	}
	v := int(g.Verbose)
	if v < 0 {
		v = 0
	}
	if v > MaxVerbose {
		v = MaxVerbose
	}
	return v
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
	Read(b []byte) (n int)
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
func (r *GReader) Read(b []byte) (n int) {
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
	Write(b []byte)
	// useful for RW paradigm
	getRollback() string
	// rollback future
	Rollback(closeBefore FilterReader) (e error)
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

func (w GWriter) Write(b []byte) {
	_, e := w.this.Write(b)
	Fatal(e)
}

// Rollback the writer and allow closing an associated reader (this can be null).
func (w GWriter) Rollback(closeBefore FilterReader) (e error) {
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
	// already handled display of errors
	return nil
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
		// already -q as command may have Notify()
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
