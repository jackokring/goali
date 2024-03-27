package main

import (
	"github.com/jackokring/goali/consts"
	goali "github.com/jackokring/goali/pkg"
)

func main() {
	// The bulk of the package utility and library code
	// Allows import of all from package goali
	// The binary name "goali" prevents naming goali
	// as a directory, so "./pkg/goali.go"
	// for the goali package
	if consts.Dynamic == "static" { // link flag for ./gob.sh

		// LDFLAGS=(
		// 	"-X '${PACKAGE}/consts.Version=${VERSION}-${COMMIT_HASH}'"
		// 	"-X '${PACKAGE}/consts.BuildTime=${BUILD_TIMESTAMP}'"
		//   )

		//   # STEP 3: Actual Go build process

		//   go generate
		//   go build -ldflags="${LDFLAGS[*]}"

		// 	"-buildmode=c-shared" go build flag.
		// Build the listed main package, plus all packages it imports,
		// into a C shared library. The only callable symbols will
		// be those functions exported using a cgo //export comment.
		// Requires exactly one main package to be listed.
		goali.Goali()
	}
}
