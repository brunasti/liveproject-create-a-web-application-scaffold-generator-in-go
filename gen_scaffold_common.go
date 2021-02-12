package main

import (
	"fmt"
)

func generateScaffoldCommon(structureType projectStructureType) int {
	logVerbose(fmt.Sprintf("Generating scaffold common content for project %s in %s", structureType.Name, structureType.Path))

	writeTemplate("handlers/handlers.go", "./"+structureType.Path+"/handlers/handlers.go", structureType)
	writeTemplate("handlers/constants.go", "./"+structureType.Path+"/handlers/constants.go", structureType)
	writeTemplate("handlers/go.mod", "./"+structureType.Path+"/handlers/go.mod", structureType)

	return 0
}
