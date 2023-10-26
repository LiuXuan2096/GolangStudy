package main

import (
	"bytes"
	"fmt"
)

func main() {
	// func1(1, 2, 3, 4, 5, 6)
	fmt.Println(printTypeValue(11, "葛诗颖", true))
}

func func1(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Print(args[i])
	}
	fmt.Println("")

	for i, data := range args {
		fmt.Println("编号为:", i)
		fmt.Println("值为:", data)
	}
}

func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer // 字节缓冲作为快速字符串拼接
	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		var typeString string
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		b.WriteString("value:")
		b.WriteString(str)
		b.WriteString(" type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}
