package main

import (
	"runtime/debug"

	"github.com/HardDie/gsplitter/cmd"
)

var Version = "dev"

func main() {
	if info, available := debug.ReadBuildInfo(); available {
		switch info.Main.Version {
		case "", "(devel)":
			// skip
		default:
			Version = info.Main.Version
		}
	}
	cmd.Execute(Version)
}
