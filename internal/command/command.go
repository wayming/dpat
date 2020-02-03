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
	otext     string
	oreceiver receiver
}

func (cmd *toUpperCmd) Execute() {
	cmd.otext = cmd.oreceiver.Action(cmd.otext)
}

type toLowerCmd struct {
	otext     string
	oreceiver receiver
}

func (cmd *toLowerCmd) Execute() {
	cmd.otext = cmd.oreceiver.Action(cmd.otext)
}

type upperStringReceiver struct {
}

func (receiver upperStringReceiver) Action(text string) string {
	return strings.ToUpper(text)
}

type lowerStringReceiver struct {
}

func (receiver lowerStringReceiver) Action(text string) string {
	return strings.ToLower(text)
}

// TestPattern test pattern
func TestPattern() {
	fmt.Println("command package")

	var upperReceiver upperStringReceiver
	testUpperCmd := &toUpperCmd{"UPPER string", upperReceiver}

	var lowerReceiver lowerStringReceiver
	testLowerCmd := &toLowerCmd{"LOWER string", lowerReceiver}

	testUpperCmd.Execute()
	testLowerCmd.Execute()
	fmt.Println(testUpperCmd.otext)
	fmt.Println(testLowerCmd.otext)
}
