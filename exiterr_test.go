package exiterr_test

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/torpidtangerine/exiterr"
)

func TestExitHandler(t *testing.T) {
	tests := []struct {
		err error

		expCode int
		expOut  string
	}{
		{
			err: nil,

			expCode: 0,
			expOut:  "",
		},
		{
			err: errors.New("unknown error"),

			expCode: 1,
			expOut:  "unknown error\n",
		},
		{
			err: context.Canceled,

			expCode: 1,
			expOut:  "",
		},
	}

	for idx, test := range tests {
		buffer := new(bytes.Buffer)
		strWriter := bufio.NewWriter(buffer)
		exitChan := make(chan int)
		eh := exiterr.NewExitHandler(strWriter, func(code int) {
			exitChan <- code
		}, exiterr.DefaultSkipOutput)

		go eh.Exit(test.err)

		code := <-exitChan
		_ = strWriter.Flush()

		if code != test.expCode {
			t.Errorf("[%d] expected '%v' received '%v'", idx, test.expCode, code)
		}

		if buffer.String() != test.expOut {
			t.Errorf("[%d] expected '%v' received '%v'", idx, test.expOut, buffer.String())
		}
	}
}
