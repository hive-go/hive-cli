package client

import (
	"github.com/hive-go/hive"
)

var ClientModule = hive.CreateModule()

func init() {
	ClientModule.AddController(ClientController)
}
