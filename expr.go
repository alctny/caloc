package main

import "math"

type Action func(a, b float64) float64

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

// 运算符优先级
var OperatorPriority = map[string]int{
	"":  -1,
	"(": 0,
	"+": 1,
	"-": 1,
	"/": 2,
	"*": 2,
	"%": 2,
	"^": 3,
}

// 运算符计算方式
var OperatorsAction = map[string]Action{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
	"%": mod,
	"^": math.Pow,
}
