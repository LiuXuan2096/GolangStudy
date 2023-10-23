package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

}

func func1() {
	cmdOutPut, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", cmdOutPut)
}
