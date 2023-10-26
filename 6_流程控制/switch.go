package main

import "fmt"

func main() {
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型~：%T", i)
	case int:
		fmt.Printf(" X 是 int 型")
	case float64:
		fmt.Printf(" x 是 float64型")
	case float32:
		fmt.Printf("x 是 float32型")
	default:
		fmt.Printf("未知型")
	}
}
