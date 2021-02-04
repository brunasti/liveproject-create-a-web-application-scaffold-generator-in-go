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

	fmt.Println("scaffol-gen application")
	fmt.Println(" - project name : ", *projectName)
	fmt.Println(" - project location : ", *projectLocation)
	fmt.Println(" - remote repository : ", *remoteRepositoryURL)
	fmt.Println(" - static assets : ", *staticAssets)
}
