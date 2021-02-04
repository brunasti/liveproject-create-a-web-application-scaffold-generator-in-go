package main

import (
	"flag"
	"fmt"
	"os"
)

type ProjectStructureType struct {
	Name          string `json:"name"`
	Path          string `json:"location"`
	RepositoryURL string `json:"repository"`
	StaticAssets  bool   `json:"static"`
}

func app_init(args []string) *ProjectStructureType {

	/*
	   -d string
	         Project location on disk
	   -n string
	         Project name
	   -r string
	         Project remote repository URL
	   -s    Project will have static assets or not
	*/

	var flagSet = flag.NewFlagSet("params", flag.ContinueOnError)
	var projectStruct = ProjectStructureType{}

	flagSet.StringVar(&projectStruct.Path, "d", "", "Project location on disk")
	flagSet.StringVar(&projectStruct.Name, "n", "", "Project name")
	flagSet.StringVar(&projectStruct.RepositoryURL, "r", "", "Project remote repository URL")
	flagSet.BoolVar(&projectStruct.StaticAssets, "s", false, "Project will have static assets or not")
	flagSet.BoolVar(&appLogVerbose, "v", false, "verbose")

	err := flagSet.Parse(args)

	if err != nil {
		// TODO Handle exit for error
		fmt.Println(err)
	}

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

	return &projectStruct
}
