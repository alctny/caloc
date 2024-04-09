package main

import (
	"fmt"
	"strings"
)

// 字符串转换为表达式栈
func Str2ExprArr(s string) ([]string, error) {
	s = strings.ReplaceAll(s, " ", "")
	expr := []string{}
	value := ""
	previousIsNumer := false
	previousChar := ' '

	for _, c := range s {
		switch c {

		case '(', ')', '^', '+', '-', '*', '/', '%':
			// 大丈夫ですか？
			if c == '-' && (previousChar == '(' || previousChar == ' ') {
				value = fmt.Sprintf("%c", c)
				continue
			} else if previousIsNumer {
				expr = append(expr, value)
				value = ""
				previousIsNumer = false
			}
			expr = append(expr, fmt.Sprintf("%c", c))

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			previousIsNumer = true
			value = fmt.Sprintf("%s%c", value, c)

		default:
			return nil, fmt.Errorf("invalid character: %c", c)
		}

		previousChar = c
	}
	if previousIsNumer {
		expr = append(expr, value)
	}
	return expr, nil
}

// 表达式预处理
func ExprPreprocess(expr string) string {
	// 去除空格
	expr = strings.ReplaceAll(expr, " ", "")
	// sin cos tan 替换为特殊字符
	//
	return expr
}
