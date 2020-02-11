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

// ToUpperCmd type
type ToUpperCmd struct {
	otext     string
	oreceiver receiver
}

// Execute ToUpperCmd
func (cmd *ToUpperCmd) Execute() {
	cmd.otext = cmd.oreceiver.Action(cmd.otext)
}

// ToLowerCmd type
type ToLowerCmd struct {
	otext     string
	oreceiver receiver
}

// Execute ToLowerCmd
func (cmd *ToLowerCmd) Execute() {
	cmd.otext = cmd.oreceiver.Action(cmd.otext)
}

type upperStringReceiver struct {
}

func (receiver upperStringReceiver) Action(text string) string {
	fmt.Println("Convert " + text + " to " + strings.ToUpper(text))
	return strings.ToUpper(text)
}

type lowerStringReceiver struct {
}

func (receiver lowerStringReceiver) Action(text string) string {
	fmt.Println("Convert " + text + " to " + strings.ToLower(text))
	return strings.ToLower(text)
}
