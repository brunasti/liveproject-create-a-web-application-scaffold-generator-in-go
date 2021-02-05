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

//func TestExitOnErrorFail(t *testing.T) {
//
//	fmt.Println("TestExitOnErrorFail : =====================")
//
//	error := true
//	err := errors.New("TEST")
//	exitOnError(err)
//	if error {
//		t.Errorf("exitOnError(err) should have killed the app")
//	}
//
//	fmt.Println("TestExitOnErrorFail : ================== end")
//}

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

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("TESTING : =======================")
	fmt.Println(m)
	fmt.Println("=================================")
	os.Exit(m.Run())
}
