package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Logger interface {
	Write(p []byte) (n int, err error)
}
type TimeLogger struct {
	timeStamp time.Time
}

func (t *TimeLogger) Write(p []byte) (int, error) {
	log.Println(t.timeStamp, string(p))
	file, err := os.OpenFile("./lv2/log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return io.WriteString(file, string(p))

}
func main() {
	logger := &TimeLogger{timeStamp: time.Now()}
	fmt.Fprintln(logger, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logger, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logger, "用户执行操作B")
}
