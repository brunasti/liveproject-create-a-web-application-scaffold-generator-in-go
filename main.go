package main

import (
	"flag"
	"fmt"
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

	fmt.Println("Generating scaffold for project", *projectName, "in", *projectLocation)
	fmt.Println("  with repository", *remoteRepositoryURL, "and static assets", *staticAssets)
}
