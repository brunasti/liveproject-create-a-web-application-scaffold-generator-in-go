package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	flagSet = setUpFlags()
	application()
}

func application() {
	fmt.Println("Scaffold Generator v 1.0 ------------")
	projectStruct = *SetConfiguration(os.Args[1:], &flagSet)

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
