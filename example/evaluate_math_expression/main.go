package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/HotPotatoC/sture/stack"
)

// operators (in order of precedence)
var operators = map[string]int{
	"^": 3, // Exponent
	"*": 2, // Multiplication
	"/": 2, // Division
	"+": 1, // Addition
	"-": 1, // Subtraction
}

func evalMathExpressionPostfix(expression string) int {
	result := stack.NewStack[int]()

	exp := strings.Split(expression, " ")

	for i := 0; i < len(exp); i++ {
		if utf8.ValidString(exp[i]) {
			if n, err := strconv.Atoi(exp[i]); err == nil {
				result.Add(n)
			}

			if _, ok := operators[exp[i]]; ok {
				y := result.Peek()
				result.Pop()
				x := result.Peek()
				result.Pop()

				switch exp[i] {
				case "+":
					result.Add(x + y)
				case "-":
					result.Add(x - y)
				case "/":
					result.Add(x / y)
				case "*":
					result.Add(x * y)
				case "^":
					result.Add(int(math.Pow(float64(x), float64(y))))
				}
			}
		}
	}

	res := result.Peek()
	result.Pop()
	return res
}

func evalMathExpressionPrefix(expression string) int {
	result := stack.NewStack[int]()

	exp := strings.Split(expression, " ")

	for i := len(exp) - 1; i >= 0; i-- {
		if utf8.ValidString(exp[i]) {
			if n, err := strconv.Atoi(exp[i]); err == nil {
				result.Add(n)
			}

			if _, ok := operators[exp[i]]; ok {
				x := result.Peek()
				result.Pop()
				y := result.Peek()
				result.Pop()

				switch exp[i] {
				case "+":
					result.Add(x + y)
				case "-":
					result.Add(x - y)
				case "/":
					result.Add(x / y)
				case "*":
					result.Add(x * y)
				case "^":
					result.Add(int(math.Pow(float64(x), float64(y))))
				}
			}
		}
	}

	res := result.Peek()
	result.Pop()
	return res
}

func main() {
	postfixExpressions := []string{
		"4 10 *",
		"4 3 * 5 +",
		"4 6 5 2 - * 3 / +",
		"7 6 5 * 3 2 ^ - +",
	}

	prefixExpressions := []string{
		"* 4 10",
		"+ 5 * 3 4",
		"+ 4 / * 6 - 5 2 3",
		"+ 7 - * 6 5 ^ 3 2",
	}

	infixExpressions := []string{
		"4 * 10",
		"5 + 3 * 4",
		"4 + 6 * ( 5 - 2 ) / 3",
		"7 + ( ( 6 * 5 ) - ( 3 ^ 2 ) )",
	}

	for _, exp := range postfixExpressions {
		result := evalMathExpressionPostfix(exp)
		fmt.Printf("%s = %d\n", exp, result)
	}

	fmt.Println()

	for _, exp := range prefixExpressions {
		result := evalMathExpressionPrefix(exp)
		fmt.Printf("%s = %d\n", exp, result)
	}

	fmt.Println("\nInfix to Postfix")

	for _, exp := range infixExpressions {
		fmt.Println(exp)
		postfixExp := infixToPostfixExpression(exp)
		result := evalMathExpressionPostfix(postfixExp)
		fmt.Printf("%s = %d\n\n", postfixExp, result)
	}

	fmt.Println("\nInfix to Prefix")

	for _, exp := range infixExpressions {
		fmt.Println(exp)
		prefixExp := infixToPrefixExpression(exp)
		result := evalMathExpressionPrefix(prefixExp)
		fmt.Printf("%s = %d\n\n", prefixExp, result)
	}
}
