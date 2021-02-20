package main

import (
	"fmt"
	"os"
	"testing"
)

var testCounter = 0
var testSuccess = 0

type generatedObject struct {
	directory bool
	found     bool
}

type generatedContent struct {
	tokens string
	found  bool
}

var generatedObjects = map[string]generatedObject{
	"handlers":              {true, false},
	"handlers/constants.go": {false, false},
	"handlers/go.mod":       {false, false},
	"handlers/handlers.go":  {false, false},
}

var generatedStaticObjects = map[string]generatedObject{
	"handlers/statics":                               {true, false},
	"handlers/statics/css":                           {true, false},
	"handlers/statics/fonts":                         {true, false},
	"handlers/statics/js":                            {true, false},
	"handlers/statics/index.html":                    {false, false},
	"handlers/statics/test.html":                     {false, false},
	"handlers/statics/css/bootstrap.min.css":         {false, false},
	"handlers/statics/css/login.css":                 {false, false},
	"handlers/statics/fonts/fontawesome-webfont.svg": {false, false},
	"handlers/statics/js/bootstrap.min.js":           {false, false},
}

var generatedContents = map[string]generatedContent{
	"handlers/constants.go": {"const appName", false},
	"handlers/go.mod":       {"go 1.16", false},
}

func checkObjects(m map[string]generatedObject, tempDir string) (bool, string) {
	objectsError := false
	errorsList := ""
	// Check if dirs & files exists
	for n, t := range m {
		fmt.Println(fmt.Sprintf("Object to check : %v - dir:%v ", n, t.directory))
		if t.directory {
			if _, err := os.Stat(tempDir + "/" + n); os.IsNotExist(err) {
				// path/to/whatever does not exist
				t.found = false
				objectsError = true
				errorsList = errorsList + fmt.Sprintf("[dir: %v] ", n)
			} else {
				t.found = true
			}
		} else {
			if _, err := os.Stat(tempDir + "/" + n); os.IsNotExist(err) {
				// path/to/whatever does not exist
				t.found = false
				objectsError = true
				errorsList = errorsList + fmt.Sprintf("[file: %v] ", n)
			} else {
				t.found = true
			}
		}
	}

	return objectsError, errorsList
}

func checkContents(m map[string]generatedContent, tempDir string) (bool, string) {
	objectsError := false
	errorsList := ""
	// Check if dirs & files exists
	for n, t := range m {
		fmt.Println(fmt.Sprintf("Content to check : %v - Content:%v ", n, t.tokens))
		if !IsExistInFile(t.tokens, tempDir+"/"+n) {
			// path/to/whatever does not exist
			t.found = false
			objectsError = true
			errorsList = errorsList + fmt.Sprintf("[file: %v - content: (%v)] ", n, t.tokens)
		} else {
			t.found = true
		}
	}

	return objectsError, errorsList
}

// --------------------------------------------------

func TestApplication(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should return errorCode 0"), func(t *testing.T) {
		tempDir := t.TempDir() + "/path"
		fmt.Println("TempDir : " + tempDir)
		ps := projectStructureType{
			Name:          "name",
			Path:          tempDir,
			RepositoryURL: "rep",
			StaticAssets:  false,
		}
		res := application(ps)
		if res != 0 {
			t.Fatalf("Expected %d return code, but got %d instead", 0, res)
		}

		objectsError, errorsList := checkObjects(generatedObjects, tempDir)

		if objectsError {
			t.Fatalf("Some of the expected objects were not created %v", errorsList)
		}

		objectsError, errorsList = checkContents(generatedContents, tempDir)

		if objectsError {
			t.Fatalf("Some of the expected contents were not found %v", errorsList)
		}

	})

	if !t.Failed() {
		testSuccess++
	}
}

func TestApplicationStatic(t *testing.T) {
	testCounter++

	t.Run(fmt.Sprintf("should return errorCode 0"), func(t *testing.T) {
		tempDir := t.TempDir() + "/path"
		fmt.Println("TempDir : " + tempDir)
		ps := projectStructureType{
			Name:          "name",
			Path:          tempDir,
			RepositoryURL: "rep",
			StaticAssets:  true,
		}
		res := application(ps)
		if res != 0 {
			t.Fatalf("Expected %d return code, but got %d instead", 0, res)
		}

		objectsError, errorsList := checkObjects(generatedObjects, tempDir)
		if objectsError {
			t.Fatalf("Some of the expected objects were not created %v", errorsList)
		}

		objectsError, errorsList = checkObjects(generatedStaticObjects, tempDir)
		if objectsError {
			t.Fatalf("Some of the expected static objects were not created %v", errorsList)
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
