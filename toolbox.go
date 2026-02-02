package main

import (
	"flag"
	"fmt"
	"runtime"
	"runtime/debug"
)

var DateBuild string

func main() {
	versionFlag := flag.Bool("version", false, "Print the version number")
	flag.Parse()
	if *versionFlag {
		if info, ok := debug.ReadBuildInfo(); ok && info != nil {
			fmt.Printf("version: %s\r\npath: %s\r\nbuild: %s\r\ngo: %s", info.Main.Version, info.Main.Path, DateBuild, info.GoVersion)
		} else {
			fmt.Printf("go: %s", runtime.Version())
		}
	}
}
