package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/utils"
)

func main() {
	utils.EnsureFileExists(config.LoggerFile)
	utils.EnsureFileExists(config.StorageFile)
	LaunchEngine()
}
