package main

import (
	"fmt"
	"time"
)

const (
	SecondsPerMinute = 60                    // 定义没分钟的秒数
	SecondPerHour    = 60 * SecondsPerMinute // 定义每小时的秒数
	SecondPerDay     = SecondPerHour * 24    // 定义每天的秒数
)

// 将传入的秒解析为3种时间单位
func resolveTime(seconds int64) (day int64, hour int64, minute int64) {
	day = seconds / SecondPerDay
	hour = seconds / SecondPerHour
	minute = seconds / SecondsPerMinute
	return
}

func main() {
	fmt.Println(resolveTime(time.Now().Unix()))
	day, hour, _ := resolveTime(time.Now().Unix()) // 只获取天和小时
	fmt.Println(day, hour)
}
