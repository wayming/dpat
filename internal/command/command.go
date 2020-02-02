package command

import (
	"fmt"
	"strings"
)

type command interface {
	Execute()
}

type receiver interface {
	Action(text string) string
}

type toUpperCmd struct {
	text     string
	receiver upperStringReceiver
}

type upperStringReceiver string

func (receiver upperStringReceiver) Action(text string) string {
	return strings.ToUpper(text)
}

func (cmd toUpperCmd) Execute() {
	cmd.text = cmd.receiver.Action(cmd.text)
}

// type toLowerCmd struct {
// 	text string
// }

// type (cmd ToLowerCmd) Execute() {
// 	text = ToLower(text)
// }

// Pprint test
func Pprint() {
	fmt.Println("command package")

	receiver := upperStringReceiver(" ")
	testUpperCmd := toUpperCmd{"abcdefg", receiver}
	testUpperCmd.Execute()
	fmt.Println(testUpperCmd.text)
}
