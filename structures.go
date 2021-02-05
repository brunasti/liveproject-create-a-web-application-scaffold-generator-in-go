package main

import (
	"flag"
)

type projectStructureType struct {
	Name          string `json:"name"`
	Path          string `json:"location"`
	RepositoryURL string `json:"repository"`
	StaticAssets  bool   `json:"static"`
}

var appLogVerbose bool
var projectStruct projectStructureType
var flagSet = flag.FlagSet{}
