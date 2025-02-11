package main

import (
    "fmt"
    "runtime/debug"
)

var (
    Version   string
    CommitHash string
    BuildTime string
)

func getVersionInfo() string {
    if info, ok := debug.ReadBuildInfo(); ok {
        for _, setting := range info.Settings {
            switch setting.Key {
            case "vcs.revision":
                CommitHash = setting.Value[:7]
            case "vcs.time":
                BuildTime = setting.Value
            }
        }
    }
    return fmt.Sprintf("Version: %s\nCommit: %s\nBuild Time: %s", 
        Version, CommitHash, BuildTime)
}

func main() {
	fmt.Println("hello world 2")
}
