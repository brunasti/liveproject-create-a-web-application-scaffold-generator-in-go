package main

import (
	"flag"
	"testing"
)

func TestAppInit(t *testing.T) {

	var args []string
	args = []string{"-d", "name"}

	// ./scaffold-gen -d ./project1 -n Project1 -r github.com/username/project1 -v

	//var projectStruct ProjectStructureType

	f := flag.NewFlagSet("params", flag.ExitOnError)
	setConfiguration(args, f)

	//projectStruct := *AppInit(args)
	//
	//if projectStruct == nil {
	//	t.Errorf("Failed")
	//}
}
