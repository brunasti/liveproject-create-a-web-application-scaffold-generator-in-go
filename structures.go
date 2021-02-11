package main

import (
	"flag"
)

type projectStructureType struct {
	Name          string `json:"name"`
	Port          string `json:"port"`
	Path          string `json:"location"`
	RepositoryURL string `json:"repository"`
	StaticAssets  bool   `json:"statics"`
}

var appLogVerbose bool
var projectStruct projectStructureType
var flagSet = flag.FlagSet{}
