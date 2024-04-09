package main

import (
	"bufio"
	"fmt"
	"os"
)

func expr(s string) (float64, error) {
	expr, err := Str2ExprArr(s)
	if err != nil {
		return 0, fmt.Errorf("表达式解析失败: %s", err)
	}
	es := NewExprStack()
	es.FromSlice(expr)
	return es.Expr(), nil
}

func run() {
	if len(os.Args) < 2 {
		input := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">")
			line, err := input.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "读取输入失败: ", err)
				os.Exit(1)
			}

			result, err := expr(string(line))
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: ", err)
				continue
			}
			fmt.Println(result)
		}
	}

	rseult, err := expr(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
	fmt.Println(rseult)
}

func main() {
	run()
}
