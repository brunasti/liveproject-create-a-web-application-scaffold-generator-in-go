package main

import "fmt"

func generateScaffoldStatic(structureType projectStructureType) int {
	if structureType.StaticAssets {
		logVerbose(fmt.Sprintf("Generating static common content for project %s in %s", structureType.Name, structureType.Path))
		copyFile("handlers/statics/index.html", "./"+structureType.Path+"/handlers/statics/index.html")
	}
	return 0
}
