package base

// RootCommand is the root command of all commands
var RootCommand *Command

func init() {
	RootCommand = &Command{
		// transform the standard command to RootCommand
		UsageLine: CommandEnv.Exec,
		Long:      "The root command",
	}
}

// RegisterCommand register a command to RootCommand
func RegisterCommand(cmd *Command) {
	RootCommand.Commands = append(RootCommand.Commands, cmd)
}
