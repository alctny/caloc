package main

import (
	"fmt"
	"math"
	"strconv"
)

type ExprStack struct {
	data []string
	top  int
}

func NewExprStack() *ExprStack {
	return &ExprStack{
		data: make([]string, 1),
		top:  -1,
	}
}

// 通过切片初始化栈
func (e *ExprStack) FromSlice(data []string) {
	e.data = make([]string, len(data))
	copy(e.data, data)
	e.top = len(data) - 1
}

// 出栈
func (e *ExprStack) Pop() string {
	if e.top >= 0 {
		e.top -= 1
		return e.data[e.top+1]
	}
	return ""
}

// 入栈
func (e *ExprStack) Push(str string) {
	e.top += 1
	if len(e.data) > e.top {
		e.data[e.top] = str
	} else {
		e.data = append(e.data, str)
	}
}

// 获取栈顶元素
func (e *ExprStack) Top() string {
	if e.top >= 0 {
		return e.data[e.top]
	}
	return ""
}

// 栈反转
func (e *ExprStack) Reverse() {
	len := e.top
	for i := 0; i <= len/2; i++ {
		e.data[i], e.data[len-i] = e.data[len-i], e.data[i]
	}
}

// 栈是否为空 true-空 false-非空
func (e *ExprStack) Empty() bool {
	return e.top == -1
}

// 将中缀表达式转化为逆波兰表达式 (Reverse Polish notation)
func (e *ExprStack) ToRPN() {
	e.Reverse()
	operator := NewExprStack()
	result := NewExprStack()

	for !e.Empty() {
		str := e.Pop()
		switch str {
		case "(":
			operator.Push(str)
		case ")":
			for operator.Top() != "(" {
				result.Push(operator.Pop())
			}
			operator.Pop()
		case "+", "-", "*", "/", "^", "%", "s", "c", "t":
			if OperationPriority[str] > OperationPriority[operator.Top()] {
				operator.Push(str)
			} else {
				for OperationPriority[str] <= OperationPriority[operator.Top()] {
					result.Push(operator.Pop())
				}
				operator.Push(str)
			}
		default:
			result.Push(str)
		}
	}
	for !operator.Empty() {
		result.Push(operator.Pop())
	}
	result.Reverse()
	e.data = result.data
	e.top = result.top
}

// 计算逆波兰表达式
func (e *ExprStack) Expr() float64 {
	parserFloat := func(s string) float64 {
		switch s {
		case "e":
			return math.E
		case "p":
			return math.Pi
		default:
			r, _ := strconv.ParseFloat(s, 64)
			return r
		}
	}

	tmp := NewExprStack()
	e.ToRPN()
	for !e.Empty() {
		str := e.Pop()
		switch str {
		case "+", "-", "*", "/", "^", "%":
			astr := tmp.Pop()
			bstr := tmp.Pop()
			a := parserFloat(astr)
			b := parserFloat(bstr)
			r := OperationAction[str](b, a)
			tmp.Push(fmt.Sprintf("%f", r))

		case "s", "c", "t":
			// TODO 三角函数操作符在操作数之前
			astr := tmp.Pop()
			a, _ := strconv.ParseFloat(astr, 64)
			r := OperationAction[str](a)
			tmp.Push(fmt.Sprintf("%f", r))
		default:
			tmp.Push(str)
		}
	}
	r, _ := strconv.ParseFloat(tmp.Top(), 64)
	return r
}
