package lab2

import (
	"fmt"
	"strings"
)

func PrefixToInfix(input string) (string, error) {
	var stack []string
	const operators = "+-*/"
	var literal = ""
	var e1 = ""
	var e2 = ""
	var res = ""

	if len(input) == 0 {
		return "", ThrowError()
	}

	var arrInput = strings.Split(input, " ")

	for i := len(arrInput) - 1; i >= 0; i-- {
		literal = arrInput[i]

		if !strings.Contains(operators, literal) {
			stack = append(stack, literal)
			continue
		}
		if len(stack) < 2 {
			return "", ThrowError()
		}

		e1, stack = Pop(stack)
		e2, stack = Pop(stack)

		e1 = SetElement(e1, literal)
		e2 = SetElement(e2, literal)

		var new = e1 + " " + literal + " " + e2
		stack = append(stack, new)
	}

	res, stack = Pop(stack)

	if len(stack) > 0 {
		return "", ThrowError()
	}

	return res, nil

}

func SetElement(e string, literal string) string {
	if strings.Contains(e, " ") && strings.Contains("*/^", literal) {
		e = "(" + e + ")"
	}
	return e
}

func Pop(array []string) (string, []string) {
	if len(array) == 0 {
		return "", []string{}
	}
	return array[len(array)-1], array[:len(array)-1]
}

func ThrowError() error {
	return fmt.Errorf("invalid input")
}
