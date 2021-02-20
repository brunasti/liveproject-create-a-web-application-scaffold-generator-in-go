package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
func createSubdir(base string, target string, header string) string {
	newpath := filepath.Join(base, target)
	logVerbose(fmt.Sprintf("%s  createSubdir - path: %s", header, newpath))
	os.MkdirAll(newpath, os.ModePerm)
	return newpath
}

// TODO test
// Copy a file
func copyFile(origin string, target string, header string) {
	logVerbose(fmt.Sprintf("%s  - Copy %s to %s", header, origin, target))

	// Open original file
	original, err := os.Open(origin)
	exitOnError(err)

	defer original.Close()

	// Create newFile file
	newFile, err := os.Create(target)
	exitOnError(err)

	defer newFile.Close()

	//This will copy
	bytesWritten, err := io.Copy(newFile, original)
	exitOnError(err)

	if bytesWritten < 0 {
		logVerbose(fmt.Sprintf("%s     Error copyng file %s : %d", header, newFile.Name(), bytesWritten))
	}
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

func getHeader(length int) string {
	var res = ""
	for i := 0; i < length; i++ {
		res = res + " "
	}
	return res
}

// TODO test
// Copy all the files in a directory
func copyDir(origin string, target string, depth int) int {
	var header = getHeader(depth * 2)
	logVerbose(fmt.Sprintf("%s  - Copy files from dir %s to dir %s", header, origin, target))

	files, err := ioutil.ReadDir(origin)
	if err != nil {
		log.Fatal(err)
	}

	var i = 0
	for _, f := range files {
		if f.IsDir() {
			//logVerbose(fmt.Sprintf("%s     Dir  : %s",header,  f.Name()))
			createSubdir(target, f.Name(), header)
			i = i + copyDir(origin+"/"+f.Name(), target+"/"+f.Name(), depth+1)

		} else {
			i++
			//logVerbose(fmt.Sprintf("%s     File : %s",header,  f.Name()))
			copyFile(origin+"/"+f.Name(), target+"/"+f.Name(), header)
		}
	}

	logVerbose(fmt.Sprintf("%s     Files copied : %d", header, i))

	return i
}

func IsExistInFile(str, filepath string) bool {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	isExist, err := regexp.Match(str, b)
	if err != nil {
		panic(err)
	}
	return isExist
}
