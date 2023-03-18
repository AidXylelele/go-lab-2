package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixToInfix(t *testing.T) {
	testCases := []struct {
		prefixExpression  string
		postfixExpression string
		err               error
	}{
		{"+ 5 * - 4 2 3", "5 + (4 - 2) * 3", nil},
		{"* 7 - ^ 3 2 * 13 4", "7 * (3 ^ 2 - 13 * 4)", nil},
		{"* / + 3 4 5 - 6 7", "((3 + 4) / 5) * (6 - 7)", nil},
		{"+ 11 + 22 + 33 + 4 + 5 + 6 + 7 8", "11 + 22 + 33 + 4 + 5 + 6 + 7 + 8", nil},
		{"- 4 3 2", "", ThrowError()},
		{"Hello world", "", ThrowError()},
		{"- - - 4", "", ThrowError()},
		{"", "", ThrowError()},
	}

	for _, tc := range testCases {
		postfixExpression, err := PrefixToInfix(tc.prefixExpression)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.postfixExpression, postfixExpression)
	}

}

func ExamplePrefixToInfix() {
	res, err := PrefixToInfix("+ 5 * - 4 2 3")
	if err == nil {
		fmt.Println(res)
	} else {
		panic(err)
	}
}
