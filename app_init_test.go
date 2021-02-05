package main

import (
	"fmt"
	"testing"
)

func TestAppInit(t *testing.T) {
	fmt.Println("TestAppInit - 01")

	var args []string
	fmt.Println("TestAppInit - 02")
	args = []string{"-d", "name"}

	fmt.Println("TestAppInit - 03", args)
	//f := flag.NewFlagSet("params", flag.ExitOnError)
	f := setUpFlags()
	fmt.Println("TestAppInit - 04", f)
	//SetConfiguration(args, f)

}
