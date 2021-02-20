package main

import (
	"fmt"
)

func createScaffoldDirs(structureType projectStructureType) int {
	logVerbose(fmt.Sprintf("Generating scaffold directories for project %s in %s", structureType.Name, structureType.Path))

	newPath := createSubdir("", structureType.Path, "")
	handlers := createSubdir(newPath, "handlers", "")

	if structureType.StaticAssets {
		createSubdir(handlers, "statics", "")
	}

	return 0
}
