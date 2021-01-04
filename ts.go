package main

import (
	"flag"
	"fmt"
	"time"
)

//显示指定日期的秒时间戳和毫秒时间戳
//根据秒时间戳和毫秒时间戳显示对应的日期
func main() {
	initFlag()
	flag.Parse()

	tz, err := time.LoadLocation(timeZone)
	if timeZone == "" || err != nil {
		fmt.Println(err)
		tz = time.Local
	}

	if inputStamp != 0 {
		fmt.Println(Stamp2Time(inputStamp))
		return
	}
	if inputTime == "" && isNow {
		inputTime = time.Now().Format(format)
	}

	print(Time2Stamp(inputTime, tz))
}

var inputTime string
var inputStamp int64
var isNow bool
var timeZone string

const format = "2006-01-02 15:04:05"

func initFlag() {
	flag.Int64Var(&inputStamp, "ts", 0, "毫秒时间戳")
	flag.StringVar(&inputTime, "dt", "", "日期时间，格式须形如2006-01-02 15:04:05")
	flag.BoolVar(&isNow, "n", true, "输出当前时间戳")
	flag.StringVar(&timeZone, "tz", time.Local.String(), "")
}
func Time2Stamp(timeStr string, loc *time.Location) (int64, int64) {
	fmt.Println("timezone is " + loc.String())
	t, err := time.ParseInLocation(format, timeStr, loc)
	if err != nil {
		fmt.Println("wrong format time!", err)
		return -1, -1
	}
	return t.In(loc).Unix(), t.In(loc).UnixNano() / 1000000
}

func Stamp2Time(ts int64) string {
	t2 := time.Unix(ts/1000, 0)
	return t2.Format("2006-01-02 15:04:05")
}

func print(tsSec, tsMicSec int64) {
	fmt.Println(inputTime)
	fmt.Println("秒", tsSec)
	fmt.Println("毫秒", tsMicSec)
}
