package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Scaffold Generator v 1.0 ------------")
	flagSet = setUpFlags()
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	application(projectStruct)
}

func application(projectStruct projectStructureType) {
	errorMessages := validateConfiguration(projectStruct)

	if len(errorMessages) > 0 {
		for i := 0; i < len(errorMessages); i++ {
			fmt.Println(errorMessages[i])
		}
		os.Exit(2)
	}

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
