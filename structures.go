package main

import (
	"flag"
)

type ProjectStructureType struct {
	Name          string `json:"name"`
	Path          string `json:"location"`
	RepositoryURL string `json:"repository"`
	StaticAssets  bool   `json:"static"`
}

var appLogVerbose bool
var projectStruct ProjectStructureType
var flagSet = flag.FlagSet{}
