package commands

import (
	"errors"

	"github.com/codegangsta/cli"
)

// CommandEnum defines each-command's argumentsSize, go-function.
type CommandEnum struct {
	// string type of command name. ex: "ls","add",...
	commandName string
	// argumentsSize defines how many arguments each-command can receive.
	argumentsSize int
	// argumentsSizeError defines error value of over argument size error.
	argumentsSizeError error
	// cmd is actual command implemented fucntion.
	cmd func(c *cli.Context)
}

// EnumValueList
var (
	LsCmd     = CommandEnum{"ls", 3, errors.New("argument size error. ls command can be used by -t ${tagName},-s ${searchWord} format. "), Ls}
	AddCmd    = CommandEnum{"add", 5, errors.New("argument size error. add command can receive bookmarkId  and up to three tags. "), Add}
	RmCmd     = CommandEnum{"rm", 2, errors.New("argument size error. rm command can receive only bookmarkId . "), Rm}
	ImportCmd = CommandEnum{"import", 2, errors.New("argument size error. import command can receive chrome'sbookmark export file. "), Import}
	OpenCmd   = CommandEnum{"open", 2, errors.New("argument size error. open command can receive only bookmarkId. "), Open}
)

// GetArgSize returns how many arguments command can receive.
func (cmdEnum *CommandEnum) GetArgSize() int {
	return cmdEnum.argumentsSize
}

// GetArgumentSizeErrorMessage returns argument-size error type
func (cmdEnum *CommandEnum) GetArgumentSizeErrorMessage() error {
	return cmdEnum.argumentsSizeError
}

// Cmd activate
func (cmdEnum *CommandEnum) Cmd(c *cli.Context) {
	cmdEnum.cmd(c)
}

// GetCommandEnum returns commandEnum type by commandname string.
func GetCommandEnum(commandName string) *CommandEnum {
	switch commandName {
	case LsCmd.commandName:
		return &LsCmd
	case AddCmd.commandName:
		return &AddCmd
	case RmCmd.commandName:
		return &RmCmd
	case ImportCmd.commandName:
		return &ImportCmd
	case OpenCmd.commandName:
		return &OpenCmd
	default:
		return nil
	}
}
