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
	skipOutput func(error) bool
}

// NewExitHandler creates an ExitHandler
func NewExitHandler(output io.StringWriter, handleExit func(int), skipOutput func(error) bool) *ExitHandler {
	return &ExitHandler{
		output:     output,
		handleExit: handleExit,
		skipOutput: skipOutput,
	}
}

// Exit calls
func (eh *ExitHandler) Exit(err error) {
	if err != nil {
		if !eh.skipOutput(err) {
			_, _ = eh.output.WriteString(err.Error() + "\n")
		}
		eh.handleExit(1)
	} else {
		eh.handleExit(0)
	}
}

// DefaultSkipOutput returns true if err == context.Canceled
func DefaultSkipOutput(err error) bool {
	return err == context.Canceled
}

// Default is the default exit handler that writes to stderr, and calls os.Exit
var Default = NewExitHandler(os.Stderr, os.Exit, DefaultSkipOutput)

// Exit calls os.Exit
var Exit = Default.Exit
