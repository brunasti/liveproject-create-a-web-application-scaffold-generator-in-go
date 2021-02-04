package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var appLogVerbose bool

func main() {

	var projectStruct ProjectStructureType

	fmt.Println("Scaffold Generator v 1.0 ------------")

	projectStruct = *app_init(os.Args[1:])

	if appLogVerbose {
		bytes, err := json.Marshal(projectStruct)
		exitOnError(err)

		fmt.Println(string(bytes))
		fmt.Println()
	}

	fmt.Println("Generating scaffold for project", projectStruct.Name, "in", projectStruct.Path)
	if appLogVerbose {
		fmt.Println("  with repository", projectStruct.RepositoryURL, "and static assets", projectStruct.StaticAssets)
	}
}
