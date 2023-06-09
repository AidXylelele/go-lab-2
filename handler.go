package lab2

import (
	"bytes"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufRead, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}
	bufRead = bytes.Trim(bufRead, "\x00")

	expression := string(bufRead)
	trimmed := strings.Trim(expression, " \n")
	res, err := PrefixToInfix(trimmed)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(res))
	return nil
}
