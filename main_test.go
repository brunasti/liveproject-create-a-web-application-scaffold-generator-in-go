package main

import (
	"fmt"
	"os"
	"testing"
)

//func TestMain(m *testing.M) {
//	// given
//	output := strings.Builder{}
//	log.SetOutput(&output)
//	const name = "test-name"
//	const path = "./test-path"
//	var expectedOutput = fmt.Sprintf("Generating scaffold for project %s in %s", name, path)
//	os.Args = []string{"scaffold_gen", "-n", name, "-d", path, "-r", "some-repo"}
//
//	// when
//	exitCode := m.Run()
//
//	// then
//	if exitCode != 0 { panic("exit code should have been 0") }
//	if output.String() != expectedOutput { panic("output is not the expected one") }
//}

func TestConfig_Validate(t *testing.T) {
	t.Run(fmt.Sprintf("should report %d errors", 3), func(t *testing.T) {
		errors := 1
		if errors != 3 {
			t.Fatalf("Expected %d errors, but got %d instead", 3, errors)
		}
	})
}

//func TestConfig_Validate(t *testing.T) {
//	testCases := []struct {
//		config projectStructureType
//		numOfErrors int
//	}{
//		{ config: projectStructureType{}, numOfErrors: 3 },
//		{ config: projectStructureType{Name: "name"}, numOfErrors: 2 },
//		{ config: projectStructureType{Name: "name", Path: "path"}, numOfErrors: 1 },
//		{ config: projectStructureType{Name: "name", Path: "path", RepositoryURL: "repository"}, numOfErrors: 0 },
//	}
//
//	for _, testCase := range testCases {
//		t.Run(fmt.Sprintf("should report %d errors", testCase.numOfErrors), func(t *testing.T) {
//			// given
//			config := testCase.config
//
//			// when
//			errors := config.Validate()
//
//			// then
//			if len(errors) != testCase.numOfErrors {
//				t.Fatalf("Expected %d errors, but got %d instead", testCase.numOfErrors, len(errors))
//			}
//		})
//	}
//}
//

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("TESTING : =======================")
	fmt.Println(m)
	fmt.Println("=================================")
	os.Exit(m.Run())
}
