package main

import (
	// "regexp"
	"testing"

	"github.com/jackokring/goali/consts"
)

// A test for package main
func TestMain(t *testing.T) {
	if consts.Dynamic != "static" && consts.Dynamic != "dynamic" {
		t.Fatalf("not compiled with valid consts.Dynamic: %s", consts.Dynamic)
	}
}
