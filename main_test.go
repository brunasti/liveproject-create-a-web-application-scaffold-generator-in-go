package main

import (
	"fmt"
	"os"
	"testing"
)

func TestExitOnError(t *testing.T) {
	fmt.Println("TestExitOnError : =======================")
	e := true
	exitOnError(nil)
	e = false
	if e {
		t.Errorf("exitOnError(nil) should have not killed the app")
	} else {
		t.Log("exitOnError(nil) Success")
	}
	fmt.Println("TestExitOnError : ================= - end")

}

func TestTest(t *testing.T) {
	errors := 3
	fmt.Println("TestTest : ==================----")
	t.Run(fmt.Sprintf("should report %d errors", 3), func(t *testing.T) {
		result := 2 + 1
		if errors != result {
			t.Fatalf("Expected %d errors, but got %d instead", 3, errors)
		}
	})
	fmt.Println("TestTest : ================== end")
}

func TestValidateConfiguration(t *testing.T) {
	fmt.Println("TestValidateConfiguration : ==================----")
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
	fmt.Println("TestValidateConfiguration : ================== end")
}

func TestValidateConfigurationFAIL(t *testing.T) {
	fmt.Println("TestValidateConfigurationFAIL : ==================----")
	t.Run(fmt.Sprintf("should report no errors"), func(t *testing.T) {
		ps := projectStructureType{
			Name:          "",
			Path:          "path",
			RepositoryURL: "rep",
			StaticAssets:  false,
		}

		res := validateConfiguration(ps)

		if len(res) == 0 || len(res) > 1 {
			t.Fatalf("Expected 1 errors, but got %d [%v]", len(res), res)
		}
	})
	fmt.Println("TestValidateConfigurationFAIL : ================== end")
}

// -----------------------------------------------------------------
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("TESTING : =======================")
	fmt.Println(m)
	fmt.Println("=================================")
	os.Exit(m.Run())
}
