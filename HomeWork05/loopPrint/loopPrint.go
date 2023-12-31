package main

import (
	"fmt"
)

func main() {
	var ch1 = make(chan struct{})
	var ch2 = make(chan struct{})
	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("这是线程A打印的", 2*i-1)
			ch1 <- struct{}{}
			ch2 <- struct{}{}
		}
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			<-ch1
			fmt.Println("这是线程B打印的", 2*i)
			<-ch2
		}
	}()
	select {}
}
