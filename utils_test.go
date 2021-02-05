package main

import (
	"testing"
)

func TestExitOnError(t *testing.T) {
	testCounter++

	e := true
	exitOnError(nil)
	e = false
	if e {
		t.Errorf("exitOnError(nil) should have not killed the app")
	}

	if !t.Failed() {
		testSuccess++
	}
}
