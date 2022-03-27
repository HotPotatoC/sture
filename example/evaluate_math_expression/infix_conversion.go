package main

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/HotPotatoC/sture/stack"
)

func infixToPostfixExpression(expression string) string {
	postfixStr := ""

	opStack := stack.NewStack[string]()

	exp := strings.Split(expression, " ")

	for i := 0; i < len(exp); i++ {
		// If current string is an operand, concat it to the postfixStr
		if _, err := strconv.Atoi(exp[i]); err == nil || unicode.IsLetter(rune(exp[i][0])) {
			postfixStr += " " + exp[i]
		} else if exp[i] == "(" { // If it's an open parentheses, push it to the stack
			opStack.Add("(")
		} else if exp[i] == ")" {
			// if it's a closing parentheses, pop the stack until an open parentheses
			// is found
			for opStack.Peek() != "(" {
				postfixStr += " " + opStack.Peek()
				opStack.Pop()
			}

			// Pop the open parentheses
			opStack.Pop()
		} else { // Otherwise it's an operator
			if exp[i] == "^" {
				// If it's an exponent, pop the stack for every operator with higher
				// precedence or equal precedence
				for !opStack.IsEmpty() && operators[exp[i]] <= operators[opStack.Peek()] {
					postfixStr += " " + opStack.Peek()
					opStack.Pop()
				}
			} else {
				// if the precedence of the operator is less than the operator in
				// the head of the stack, Pop all the operators in the stack.
				for !opStack.IsEmpty() && operators[exp[i]] < operators[opStack.Peek()] {
					postfixStr += " " + opStack.Peek()
					opStack.Pop()
				}
			}

			// Now, push the current operator to the stack
			opStack.Add(exp[i])
		}
	}

	// Push all the remaining operators to the postfixStr
	for !opStack.IsEmpty() {
		postfixStr += " " + opStack.Peek()
		opStack.Pop()
	}

	return postfixStr[1:] // Remove the first space
}

func infixToPrefixExpression(expression string) string {
	return stringReverse(infixToPostfixExpression(stringReverse(expression)))
}

func stringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// if rune is '(' replace it with ')' and vice versa
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			runes[i] = ')'
		} else if runes[i] == ')' {
			runes[i] = '('
		}
	}

	return string(runes)
}
