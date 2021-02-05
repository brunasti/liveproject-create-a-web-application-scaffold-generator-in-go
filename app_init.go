package main

import (
	"flag"
	"fmt"
	"os"
)

func setUpFlags() flag.FlagSet {
	/*
	   -d string
	         Project location on disk
	   -n string
	         Project name
	   -r string
	         Project remote repository URL
	   -s    Project will have static assets or not
	*/

	var appFlagSet = flag.NewFlagSet("params", flag.ContinueOnError)

	appFlagSet.StringVar(&projectStruct.Path, "d", "", "Project location on disk")
	appFlagSet.StringVar(&projectStruct.Name, "n", "", "Project name")
	appFlagSet.StringVar(&projectStruct.RepositoryURL, "r", "", "Project remote repository URL")
	appFlagSet.BoolVar(&projectStruct.StaticAssets, "s", false, "Project will have static assets or not")
	appFlagSet.BoolVar(&appLogVerbose, "v", false, "verbose")

	return *appFlagSet
}

func setConfiguration(args []string, flagSet *flag.FlagSet) *ProjectStructureType {

	err := flagSet.Parse(args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
