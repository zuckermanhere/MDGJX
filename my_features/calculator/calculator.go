package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 定义命令行参数
	a := flag.Float64("a", 0, "第一个数字")
	b := flag.Float64("b", 0, "第二个数字")
	op := flag.String("op", "+", "运算符: +, -, *, /")
	flag.Parse()

	// 计算结果
	var result float64
	switch *op {
	case "+":
		result = *a + *b
	case "-":
		result = *a - *b
	case "*":
		result = *a * *b
	case "/":
		if *b == 0 {
			fmt.Println("错误：除数不能为 0")
			os.Exit(1)
		}
		result = *a / *b
	default:
		fmt.Printf("错误：不支持的运算符 '%s'（支持的运算符: +, -, *, /）\n", *op)
		os.Exit(1)
	}

	// 格式化输出，小数无意义时显示整数
	aStr := formatNumber(*a)
	bStr := formatNumber(*b)
	rStr := formatNumber(result)

	fmt.Printf("%s %s %s = %s\n", aStr, *op, bStr, rStr)
}

// formatNumber 将浮点数格式化为简洁字符串，去掉多余的 .0
func formatNumber(n float64) string {
	if n == float64(int64(n)) {
		return strconv.FormatInt(int64(n), 10)
	}
	return strconv.FormatFloat(n, 'f', 6, 64)
}