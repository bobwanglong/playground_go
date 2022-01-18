package main

import (
	"log"
	"time"
)

func main() {
	// old := time.Now()
	// time.Sleep(time.Second)
	// // duration := old.Sub(time.Now())
	// duration := time.Until(old) // t.sub的缩写，表示过去一段时间持续了多久

	// fmt.Println(duration)
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStamp := time.Now().Unix()
	formatTimeStr := time.Unix(timeStamp, 0).Format("2006-01-02 15:04")
	formatTimeStr = formatTimeStr + ":00"
	// formatTime, _ := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	stamp, _ := time.ParseInLocation(timeTemplate1, formatTimeStr, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())
	// tt := formatTime.Unix()
	// formatTimeStr2 := time.Unix(tt, 0).Format("2006-01-02 15:04:05")
	// fmt.Println(formatTimeStr, formatTimeStr2)

}
