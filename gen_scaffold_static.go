package main

import "fmt"

func generateScaffoldStatic(structureType projectStructureType) int {
	if structureType.StaticAssets {
		logVerbose(fmt.Sprintf("Generating static common content for project %s in %s", structureType.Name, structureType.Path))
		copyDir("handlers/statics", structureType.Path+"/handlers/statics", 0)
	}
	return 0
}
