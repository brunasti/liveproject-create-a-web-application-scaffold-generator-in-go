package main

import (
	"encoding/json"
	"fmt"
)

var appLogVerbose bool

func main() {

	var projectStruct projectStructureType

	fmt.Println("Scaffold Generator v 1.0 ------------")

	projectStruct = app_init()

	if appLogVerbose {
		bytes, err := json.Marshal(projectStruct)
		exitOnError(err)

		fmt.Println(string(bytes))
		fmt.Println()
	}

	fmt.Println("Generating scaffold for project", projectStruct.Name, "in", projectStruct.Path)
	fmt.Println("  with repository", projectStruct.RepositoryURL, "and static assets", projectStruct.StaticAssets)
}
