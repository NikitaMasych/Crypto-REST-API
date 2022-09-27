package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/platform"
)

func main() {
	platform.EnsureFileExists(config.Get().LoggerFile)
	platform.EnsureFileExists(config.Get().StorageFile)
	LaunchEngine()
}
