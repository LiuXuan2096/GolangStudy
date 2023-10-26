package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 字符串切片()
	字符串循环()
}

func 字符串切片() {
	var cmd = "notepad123"
	fmt.Printf("你要执行的命令%s\n", cmd[0:7])
	exec.Command(cmd[0:7]).Run()
}

func 字符串循环() {
	var cmd = "葛诗颖"
	fmt.Println("----------------------------------")
	for _, ch := range cmd {
		fmt.Printf("%c: %d\n", ch, ch)
	}
}
