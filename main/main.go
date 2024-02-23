package main

import (
	"flag"
	"os"

	"github.com/lmmqxyx404/my_core/main/commands/base"
	// 必须加载这个 package
	_ "github.com/lmmqxyx404/my_core/main/distro/all"
)

// 程序入口
func main() {
	// 生成可以被解析的命令
	os.Args = getArgsV4Compatible()
	base.RootCommand.Long = "Xrays is a platform for building proxies."

	base.RootCommand.Commands = append(
		[]*base.Command{
			cmdRun,
			cmdVersion,
		},
		base.RootCommand.Commands...,
	)
	base.Execute()
}

// 主要处理传进来的 命令行参数
func getArgsV4Compatible() []string {
	// 1. 默认执行 run 参数
	if len(os.Args) == 1 {
		return []string{os.Args[0], "run"}
	}
	// 2. 处理 '-'
	if os.Args[1][0] != '-' {
		return os.Args
	}
	version := false
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.BoolVar(&version, "version", false, "")
	// parse silently, no usage, no error output
	fs.Usage = func() {}

	fs.SetOutput(&null{})
	err := fs.Parse(os.Args[1:])
	if err == flag.ErrHelp {
		// fmt.Println("DEPRECATED: -h, WILL BE REMOVED IN V5.")
		// fmt.Println("PLEASE USE: xray help")
		// fmt.Println()
		return []string{os.Args[0], "help"}
	}
	if version {
		// fmt.Println("DEPRECATED: -version, WILL BE REMOVED IN V5.")
		// fmt.Println("PLEASE USE: xray version")
		// fmt.Println()
		return []string{os.Args[0], "version"}
	}
	// fmt.Println("COMPATIBLE MODE, DEPRECATED.")
	// fmt.Println("PLEASE USE: xray run [arguments] INSTEAD.")
	// fmt.Println()
	return append([]string{os.Args[0], "run"}, os.Args[1:]...)
}

type null struct{}

func (n *null) Write(p []byte) (int, error) {
	return len(p), nil
}
