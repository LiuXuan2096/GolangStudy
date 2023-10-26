package main

import (
	"fmt"
	"time"
)

func useTimes(myfunc func(int) int, arg int) {
	startTime := time.Now()
	myfunc(arg)
	fmt.Println(time.Since(startTime))
}

func main() {
	useTimes(func(arg int) (sum int) {
		for i := 0; i < arg; i++ {
			sum += i
		}
		return sum
	}, 10000000000)
}
