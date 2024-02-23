package main

import (
	"fmt"

	"github.com/lmmqxyx404/my_core/core"
	"github.com/lmmqxyx404/my_core/main/commands/base"
)

var cmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "Show current version of Xray",
	Long: `Version prints the build information for Xray executables.
	`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

// must use the core description
func printVersion() {
	// anchor1: call the xray core
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
