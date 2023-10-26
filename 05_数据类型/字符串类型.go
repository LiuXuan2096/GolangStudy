package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cmd string
	fmt.Scanf("%s", &cmd)
	fmt.Printf("你要执行的命令是%s\n", cmd)
	exec.Command(cmd).Run()
}
