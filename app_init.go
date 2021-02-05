package main

import (
	"flag"
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

	var appFlagSet = flag.NewFlagSet(appName, flag.ContinueOnError)

	appFlagSet.StringVar(&projectStruct.Path, "d", "", "Project location on disk")
	appFlagSet.StringVar(&projectStruct.Name, "n", "", "Project name")
	appFlagSet.StringVar(&projectStruct.RepositoryURL, "r", "", "Project remote repository URL")
	appFlagSet.BoolVar(&projectStruct.StaticAssets, "s", false, "Project will have static assets or not")
	appFlagSet.BoolVar(&appLogVerbose, "v", false, "verbose")

	return *appFlagSet
}

func validateConfiguration(projectStruct projectStructureType) []string {
	var errorMessages []string
	var message string

	if projectStruct.Name == "" {
		message = "Project name cannot be empty"
		errorMessages = append(errorMessages, message)
	}

	if projectStruct.Path == "" {
		message = "Project path cannot be empty"
		errorMessages = append(errorMessages, message)
	}

	if projectStruct.RepositoryURL == "" {
		message = "Project repository URL cannot be empty"
		errorMessages = append(errorMessages, message)
	}

	return errorMessages
}
