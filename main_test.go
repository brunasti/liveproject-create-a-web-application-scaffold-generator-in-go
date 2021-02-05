package main

import (
	"fmt"
	"os"
	"testing"
)

var testCounter = 0
var testSuccess = 0

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

func TestTest(t *testing.T) {
	testCounter++

	errors := 3
	t.Run(fmt.Sprintf("should report %d errors", 3), func(t *testing.T) {
		result := 2 + 1
		if errors != result {
			t.Fatalf("Expected %d errors, but got %d instead", 3, errors)
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

// -----------------------------------------------------------------
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("=================================")
	fmt.Println("Scaffold Generator Test suite ---")
	//fmt.Println("TESTING : =======================")
	//fmt.Println(*m)
	fmt.Println("=================================")

	// THIS IS THE CORE OF THE TEST EXECUTION
	res := m.Run()

	fmt.Println("=================================")
	fmt.Println("Scaffold Generator Test suite res")
	fmt.Println("=================================")
	fmt.Println("Errors reported", res)
	fmt.Println("=================================")
	fmt.Printf("Executed %3d tests", testCounter)
	fmt.Println()
	fmt.Printf("Success  %3d tests", testSuccess)
	fmt.Println()
	fmt.Printf("FAILED   %3d tests", testCounter-testSuccess)
	fmt.Println()

	os.Exit(res)
}
