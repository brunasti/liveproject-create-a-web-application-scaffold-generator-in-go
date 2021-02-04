package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var appLogVerbose bool

func main() {

	//type ProjectStructureType struct {
	//	Name          string
	//	Path          string
	//	RepositoryURL string
	//	StaticAssets  bool
	//}
	//
	//var projectStruct ProjectStructureType

	/*
	   -d string
	         Project location on disk
	   -n string
	         Project name
	   -r string
	         Project remote repository URL
	   -s    Project will have static assets or not
	*/

	//type projectStructureType struct {
	//	Name          string
	//	Path          string
	//	RepositoryURL string
	//	StaticAssets  bool
	//}

	var projectStruct projectStructureType

	//projectStructPath := flag.String("d", "", "Project location on disk")
	//projectStructName := flag.String("n", "", "Project name")
	//projectStructRepositoryURL := flag.String("r", "", "Project remote repository URL")
	//projectStructStaticAssets := flag.Bool("s", false, "Project will have static assets or not")
	//
	//flag.Parse()
	//
	//fmt.Println("Name ", projectStructName)
	//
	//projectStruct.Path = *projectStructPath
	//projectStruct.Name = *projectStructName
	//projectStruct.RepositoryURL = *projectStructRepositoryURL
	//projectStruct.StaticAssets = *projectStructStaticAssets
	//
	//fmt.Println("ARGS ", flag.Args())
	//
	//
	//flagError := false
	//
	//if projectStruct.Name == "" {
	//	fmt.Println("Project name cannot be empty")
	//	flagError = true
	//}
	//
	//if projectStruct.Path == "" {
	//	fmt.Println("Project path cannot be empty")
	//	flagError = true
	//}
	//
	//if projectStruct.RepositoryURL == "" {
	//	fmt.Println("Project repository URL cannot be empty")
	//	flagError = true
	//}

	projectStruct = app_init()

	if appLogVerbose {
		fmt.Println("------------")
		fmt.Println(projectStruct)
		fmt.Println("------------")
		bytes, err := json.Marshal(projectStruct)
		exitOnError(err)

		fmt.Println(string(bytes))
		fmt.Println("------------")
	}

	fmt.Println("Generating scaffold for project", projectStruct.Name, "in", projectStruct.Path)
	fmt.Println("  with repository", projectStruct.RepositoryURL, "and static assets", projectStruct.StaticAssets)
}

// exitOnError prints any errors and exits.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
