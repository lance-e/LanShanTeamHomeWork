package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	path := "./lv1/test.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	message := "helloworld"
	t1 := time.Now()
	for i := 0; i <= 100; i++ {
		io.WriteString(file, message)
	}
	duration1 := time.Since(t1)
	writer := bufio.NewWriter(file)
	t2 := time.Now()
	for i := 0; i <= 100; i++ {
		writer.WriteString(message)
	}
	writer.Flush()
	duration2 := time.Since(t2)
	fmt.Println("不带缓存的io操作：", duration1)
	fmt.Println("带缓存的io操作：", duration2)
}
