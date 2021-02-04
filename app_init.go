package main

import (
	"flag"
	"fmt"
	"os"
)

type projectStructureType struct {
	Name          string
	Path          string
	RepositoryURL string
	StaticAssets  bool
}

func app_init() projectStructureType {

	/*
	   -d string
	         Project location on disk
	   -n string
	         Project name
	   -r string
	         Project remote repository URL
	   -s    Project will have static assets or not
	*/

	var projectStruct projectStructureType

	projectStructPath := flag.String("d", "", "Project location on disk")
	projectStructName := flag.String("n", "", "Project name")
	projectStructRepositoryURL := flag.String("r", "", "Project remote repository URL")
	projectStructStaticAssets := flag.Bool("s", false, "Project will have static assets or not")
	verbose := flag.Bool("v", false, "verbose")

	flag.Parse()

	projectStruct.Path = *projectStructPath
	projectStruct.Name = *projectStructName
	projectStruct.RepositoryURL = *projectStructRepositoryURL
	projectStruct.StaticAssets = *projectStructStaticAssets
	appLogVerbose = *verbose

	flagError := false

	if projectStruct.Name == "" {
		fmt.Println("Project name cannot be empty")
		flagError = true
	}

	if projectStruct.Path == "" {
		fmt.Println("Project path cannot be empty")
		flagError = true
	}

	if projectStruct.RepositoryURL == "" {
		fmt.Println("Project repository URL cannot be empty")
		flagError = true
	}

	if flagError {
		os.Exit(2)
	}

	return projectStruct
}
