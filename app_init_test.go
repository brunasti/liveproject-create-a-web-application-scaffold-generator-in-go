package main

import (
	"fmt"
	"testing"
)

func TestSetUpFlags(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("no errors"), func(t *testing.T) {
		var res = setUpFlags()

		if appName != res.Name() {
			t.Fatalf("Expected flag.FlagSet name as [%v], got [%v]", appName, res.Name())
		}

		if res.Lookup("n") == nil {
			t.Fatalf("Name parameter not defined")
		}

		if res.Lookup("x") != nil {
			t.Fatalf("Strange parameter defined !!!")
		}
	})

	if !t.Failed() {
		testSuccess++
	}

}

func TestSetUpFlagsParse(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("parsing the command parameters"), func(t *testing.T) {
		var flagSet = setUpFlags()

		var input = []string{"-n", "Name", "-d", "Path", "-r", "Repo", "-s", "false", "-v", "true"}

		err := flagSet.Parse(input)

		if err != nil {
			t.Fatalf("Error parsing the parameters %v", err)
		}

		if !flagSet.Parsed() {
			t.Fatalf("Error parsing the parameters %v", input)
		}

		if projectStruct.Name != input[1] {
			t.Fatalf("Error for parameter [%v] - expected %v but got %v", input[0], input[1], projectStruct.Name)
		}

		if projectStruct.Path != input[3] {
			t.Fatalf("Error for parameter [%v] - expected %v but got %v", input[2], input[3], projectStruct.Path)
		}

		if projectStruct.RepositoryURL != input[5] {
			t.Fatalf("Error for parameter [%v] - expected %v but got %v", input[4], input[5], projectStruct.RepositoryURL)
		}

		if !projectStruct.StaticAssets {
			t.Fatalf("Error for parameter [%v] - expected %v but got %v", input[6], input[7], projectStruct.StaticAssets)
		}

		if appLogVerbose {
			t.Fatalf("Error for parameter [%v] - expected %v but got %v", input[8], input[9], appLogVerbose)
		}

	})

	if !t.Failed() {
		testSuccess++
	}

}

func TestSetUpFlagsParseHelp(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("parsing the command parameters"), func(t *testing.T) {
		var flagSet = setUpFlags()

		var input = []string{"-h"}

		err := flagSet.Parse(input)

		if !flagSet.Parsed() {
			t.Fatalf("Error parsing the parameters %v", input)
		}

		if err != nil {
			var serr = fmt.Sprintf("%v", err)
			if serr != "flag: help requested" {
				t.Fatalf("Error parsing the command parameters -h : %v", err)
			}
		} else {
			t.Fatalf("Error parsing the command parameters -h")
		}

	})

	if !t.Failed() {
		testSuccess++
	}

}

func TestValidateConfiguration(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should report no errors"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "name",
			Path:          "path",
			RepositoryURL: "rep",
			StaticAssets:  false,
		}

		res := validateConfiguration(ps)

		if len(res) > 0 {
			t.Fatalf("Expected 0 errors, but got %d [%v]", len(res), res)
		}
	})

	if !t.Failed() {
		testSuccess++
	}
}

func TestValidateConfigurationFAIL(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should report 1 errors"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "",
			Path:          "path",
			RepositoryURL: "rep",
			StaticAssets:  false,
		}

		res := validateConfiguration(ps)

		if len(res) != 1 {
			t.Fatalf("Expected 1 errors, but got %d [%v]", len(res), res)
		}
	})

	if !t.Failed() {
		testSuccess++
	}

}

func TestValidateConfigurationFAIL2(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should report 2 errors"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "",
			Path:          "path",
			RepositoryURL: "",
			StaticAssets:  true,
		}

		res := validateConfiguration(ps)

		if len(res) != 2 {
			t.Fatalf("Expected 2 errors, but got %d [%v]", len(res), res)
		}
	})

	if !t.Failed() {
		testSuccess++
	}
}

func TestValidateConfigurationFAIL3(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should report 3 errors"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "",
			Path:          "",
			RepositoryURL: "",
			StaticAssets:  true,
		}

		res := validateConfiguration(ps)

		if len(res) != 3 {
			t.Fatalf("Expected 3 errors, but got %d [%v]", len(res), res)
		}
	})

	if !t.Failed() {
		testSuccess++
	}
}
