package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestComputeHandler(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("- 6 2"),
		Output: b,
	}
	err := handler.Compute()

	assert.Equal(t, err, nil)
	assert.Equal(t, b.String(), "6 - 2")
}

func TestComputeHandlerHard(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("+ 5 - 6 7"),
		Output: b,
	}
	err := handler.Compute()

	assert.Equal(t, err, nil)
	assert.Equal(t, b.String(), "5 + 6 - 7")
}

func TestComputeHandlerError(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("14 88"),
		Output: b,
	}
	err := handler.Compute()

	assert.Equal(t, err, ThrowError())
}
