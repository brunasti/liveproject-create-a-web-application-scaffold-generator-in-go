package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	/*
	   -d string
	         Project location on disk
	   -n string
	         Project name
	   -r string
	         Project remote repository URL
	   -s    Project will have static assets or not
	*/

	type projectStructureType struct {
		Name          string
		Path          string
		RepositoryURL string
		StaticAssets  bool
	}

	projectLocation := flag.String("d", "", "Project location on disk")
	projectName := flag.String("n", "", "Project name")
	remoteRepositoryURL := flag.String("r", "", "Project remote repository URL")
	staticAssets := flag.Bool("s", false, "Project will have static assets or not")

	flag.Parse()

	flagError := false

	if *projectName == "" {
		fmt.Println("Project name cannot be empty")
		flagError = true
	}

	if *projectLocation == "" {
		fmt.Println("Project path cannot be empty")
		flagError = true
	}

	if *remoteRepositoryURL == "" {
		fmt.Println("Project repository URL cannot be empty")
		flagError = true
	}

	if flagError {
		return
	}

	var projectStruct projectStructureType

	projectStruct.Name = *projectName
	projectStruct.Path = *projectLocation
	projectStruct.RepositoryURL = *remoteRepositoryURL
	projectStruct.StaticAssets = *staticAssets

	fmt.Println("------------")
	fmt.Println(projectStruct)
	fmt.Println("------------")
	bytes, err := json.Marshal(projectStruct)
	exitOnError(err)

	fmt.Println(string(bytes))
	fmt.Println("------------")

	fmt.Println("Generating scaffold for project", projectStruct.Name, "in", *projectLocation)
	fmt.Println("  with repository", *remoteRepositoryURL, "and static assets", *staticAssets)
}

// exitOnError prints any errors and exits.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
