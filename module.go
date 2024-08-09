package main

var ModuleString = `package #MODULE_NAME

import (
	"github.com/hive-go/hive"
)

var #CAPITALIZED_NAMEModule = hive.CreateModule()

func init() {
	#CAPITALIZED_NAMEModule.AddController(#CAPITALIZED_NAMEController)
}
`
