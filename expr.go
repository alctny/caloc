package main

import (
	"fmt"
	"math"
	"strings"
)

type Action func(...float64) float64

func Binaryperation2Action(f func(float64, float64) float64) Action {
	return func(nums ...float64) float64 {
		if len(nums) < 2 {
			return math.NaN()
		}
		return f(nums[0], nums[1])
	}
}

func UnaryOperation2Action(f func(float64) float64) Action {
	return func(nums ...float64) float64 {
		if len(nums) < 1 {
			return math.NaN()
		}
		return f(nums[0] * math.Pi / 180.0)
	}
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func div(a, b float64) float64 {
	return a / b
}

func mul(a, b float64) float64 {
	return a * b
}

func mod(a, b float64) float64 {
	return float64(int(a) % int(b))
}

func log(a, b float64) float64 {
	return math.Log(b) / math.Log(a)
}

// 运算符优先级
var OperationPriority = map[string]int{
	"":  -1,
	"(": 0,
	"+": 1,
	"-": 1,
	"/": 2,
	"*": 2,
	"%": 2,
	"^": 3,

	"s": 4,
	"c": 4,
	"t": 4,
	"l": 4,
	"g": 4,
	"n": 4,
}

// 运算符计算方式
var OperationAction = map[string]Action{
	"+": Binaryperation2Action(add),
	"-": Binaryperation2Action(sub),
	"*": Binaryperation2Action(mul),
	"/": Binaryperation2Action(div),
	"%": Binaryperation2Action(mod),
	"^": Binaryperation2Action(math.Pow),
	"l": Binaryperation2Action(log),

	"s": UnaryOperation2Action(math.Sin),
	"c": UnaryOperation2Action(math.Cos),
	"t": UnaryOperation2Action(math.Tan),
	"g": UnaryOperation2Action(math.Log10),
	"n": UnaryOperation2Action(math.Log),
}

// 字符串转换为表达式栈
func Str2ExprArr(str string) ([]string, error) {
	str = ExprPreprocess(str)

	expr := []string{}
	number := ""
	previousIsNumer := false
	previousChar := ' '

	for _, c := range str {
		switch c {

		case '(', ')', '^', '+', '-', '*', '/', '%', 's', 'c', 't':
			// 如果 - 号之前为左括号或表达式以 - 开头说明 - 为负号
			if c == '-' && (previousChar == '(' || previousChar == ' ') {
				number = fmt.Sprintf("%c", c)
				continue
			} else if previousIsNumer {
				expr = append(expr, number)
				number = ""
				previousIsNumer = false
			}
			expr = append(expr, fmt.Sprintf("%c", c))

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			previousIsNumer = true
			number = fmt.Sprintf("%s%c", number, c)

		case 'e', 'p', 'i':
			expr = append(expr, fmt.Sprintf("%c", c))

		default:
			return nil, fmt.Errorf("invalid character: %c", c)
		}

		previousChar = c
	}
	if previousIsNumer {
		expr = append(expr, number)
	}
	return expr, nil
}

// 表达式预处理
func ExprPreprocess(expr string) string {
	// 去除空白符
	expr = strings.ReplaceAll(expr, " ", "")
	expr = strings.ReplaceAll(expr, "\n", "")
	expr = strings.ReplaceAll(expr, "\t", "")
	// sin cos tan 替换为特殊字符
	expr = strings.ReplaceAll(expr, "sin", "s")
	expr = strings.ReplaceAll(expr, "cos", "c")
	expr = strings.ReplaceAll(expr, "tan", "t")
	expr = strings.ReplaceAll(expr, "log", "l")
	expr = strings.ReplaceAll(expr, "lg", "g")
	expr = strings.ReplaceAll(expr, "ln", "n")

	return expr
}
