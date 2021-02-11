package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// exitOnError prints any errors and exits.
// TODO Add message and exit code
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO test
func logVerbose(msg string) {
	if appLogVerbose {
		fmt.Println(msg)
	}
}

// TODO test
func createSubdir(base string, new string) string {
	newpath := filepath.Join(base, new)
	logVerbose(fmt.Sprintf("  - path: %s", newpath))
	os.MkdirAll(newpath, os.ModePerm)
	return newpath
}

// TODO test
// Copy a file
func copyFile(origin string, target string) {
	logVerbose(fmt.Sprintf("  - Copy %s to %s", origin, target))

	// Open original file
	original, err := os.Open(origin)
	exitOnError(err)
	logVerbose(fmt.Sprintf("     Opened %s", origin))

	defer original.Close()

	// Create newFile file
	newFile, err := os.Create(target)
	exitOnError(err)
	logVerbose(fmt.Sprintf("     Opened %s", newFile.Name()))

	defer newFile.Close()

	//This will copy
	bytesWritten, err := io.Copy(newFile, original)
	exitOnError(err)

	logVerbose(fmt.Sprintf("     Bytes Written: %d", bytesWritten))
}

// TODO test
// Write into a file
func writeFile(target string, content string) {
	logVerbose(fmt.Sprintf("  - Write %s", target))

	f, err := os.Create(target)
	exitOnError(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	bytesWritten, err := w.WriteString(content)
	exitOnError(err)

	w.Flush()
	f.Close()

	logVerbose(fmt.Sprintf("     Bytes Written: %d", bytesWritten))
}

// TODO test
// Write a template into a file
func writeTemplate(origin string, target string, content projectStructureType) {
	logVerbose(fmt.Sprintf("  - Template %s - Write %s", origin, target))

	f, err := os.Create(target)
	exitOnError(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	t, _ := template.ParseFiles(origin)
	err = t.Execute(w, content)
	exitOnError(err)

	w.Flush()
	f.Close()
}
