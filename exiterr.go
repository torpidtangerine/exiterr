package exiterr

import (
	"context"
	"io"
	"os"
)

// ExitHandler handles process exit calls
type ExitHandler struct {
	output     io.StringWriter
	handleExit func(int)
}

// NewExitHandler creates an ExitHandler
func NewExitHandler(output io.StringWriter, handleExit func(int)) *ExitHandler {
	return &ExitHandler{
		output:     output,
		handleExit: handleExit,
	}
}

// Exit calls
func (eh *ExitHandler) Exit(err error) {
	if err != nil {
		if err != context.Canceled {
			_, _ = eh.output.WriteString(err.Error())
			_, _ = eh.output.WriteString("\n")
		}
		eh.handleExit(1)
	} else {
		eh.handleExit(0)
	}
}

// Default is the default exit handler that writes to stderr, and calls os.Exit
var Default = NewExitHandler(os.Stderr, os.Exit)

// Exit calls os.Exit
var Exit = Default.Exit
