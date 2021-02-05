package main

import (
	"fmt"
	"os"
	"testing"
)

var testCounter = 0
var testSuccess = 0

func TestApplication(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should return errorCode 0"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "name",
			Path:          "path",
			RepositoryURL: "rep",
			StaticAssets:  false,
		}
		res := application(ps)
		if res != 0 {
			t.Fatalf("Expected %d return code, but got %d instead", 0, res)
		}
	})

	if !t.Failed() {
		testSuccess++
	}
}

func TestApplicationFail(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should return errorCode -2"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "",
			Path:          "path",
			RepositoryURL: "",
			StaticAssets:  false,
		}
		res := application(ps)
		if res != -2 {
			t.Fatalf("Expected %d return code, but got %d instead", -2, res)
		}
	})

	if !t.Failed() {
		testSuccess++
	}
}

// -----------------------------------------------------------------
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("========================================")
	fmt.Println(appName, " Test suite ---")
	fmt.Println("========================================")

	// THIS IS THE CORE OF THE TEST EXECUTION
	res := m.Run()

	fmt.Println("========================================")
	fmt.Println(appName, " Test suite res")
	fmt.Println("========================================")
	fmt.Println("Errors reported", res)
	fmt.Println("========================================")
	fmt.Printf("Executed %3d tests", testCounter)
	fmt.Println()
	fmt.Printf("Success  %3d tests", testSuccess)
	fmt.Println()
	fmt.Printf("FAILED   %3d tests", testCounter-testSuccess)
	fmt.Println()

	os.Exit(res)
}
