package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println(appName + " ------------")
	flagSet = setUpFlags()

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	res := application(projectStruct)
	if res < 0 {
		os.Exit(-res)
	}
}

func application(projectStruct projectStructureType) int {
	errorMessages := validateConfiguration(projectStruct)

	if len(errorMessages) > 0 {
		for i := 0; i < len(errorMessages); i++ {
			fmt.Println(errorMessages[i])
		}
		return -2
	}

	if appLogVerbose {
		bytes, err := json.Marshal(projectStruct)
		exitOnError(err)

		fmt.Println(string(bytes))
		fmt.Println()
	}

	fmt.Println("Generating scaffold for project", projectStruct.Name, "in", projectStruct.Path)
	logVerbose(fmt.Sprintf("  with repository %s and statics assets %v", projectStruct.RepositoryURL, projectStruct.StaticAssets))

	res := createScaffoldDirs(projectStruct)
	res = generateScaffoldCommon(projectStruct)
	res = generateScaffoldStatic(projectStruct)

	return res
}
