package app

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/service/commands"
)

//ArgumentSizeValidator validate command argument size.
func ArgumentSizeValidator(c *cli.Context) error {
	var commandName = c.Args().Get(0)
	var cmdEnum = commands.GetCommandEnum(commandName)
	if cmdEnum == nil {
		return nil
	}
	if len(c.Args()) > cmdEnum.GetArgSize() {
		return cmdEnum.GetArgumentSizeErrorMessage()
	}
	return nil
}
