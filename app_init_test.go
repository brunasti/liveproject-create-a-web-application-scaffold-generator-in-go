package main

import (
	"fmt"
	"testing"
)

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
