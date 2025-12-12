package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func DoExpression(expression string) (float64, error) {
	numbera, numberb, operator, err := parseExpression(expression)
	if err != nil {
		return 0, err
	}
	var result float64
	switch operator {
	case "+":
		result = numbera + numberb
	case "-":
		result = numbera - numberb
	case "*":
		result = numbera * numberb
	case "/":
		result = numbera / numberb
	default:
		return 0, fmt.Errorf("invalid operator")
	}
	return result, nil
}

func parseExpression(expression string) (numbera float64, numberb float64, operator string, err error) {
	// 输入 10+20 根据中间的运算符进行分割
	for i, char := range expression {
		if char == '+' || char == '-' || char == '*' || char == '/' {
			numbera, err = strconv.ParseFloat(strings.TrimSpace(expression[:i]), 64)
			if err != nil {
				return 0, 0, "", err
			}
			numberb, err = strconv.ParseFloat(strings.TrimSpace(expression[i+1:]), 64)
			if err != nil {
				return 0, 0, "", err
			}
			operator = string(char)
			return numbera, numberb, operator, nil
		}
	}

	return 0, 0, "", nil
}
