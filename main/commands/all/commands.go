package all

import (
	"github.com/lmmqxyx404/my_core/main/commands/base"
)

// go:generate go run github.com/xtls/xray-core/common/errors/errorgen

func init() {

	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		// todo: add more commands
		cmdUUID,
	)
}
