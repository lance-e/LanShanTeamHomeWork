package main

import (
	"LanShanHomeWork2/model"
	"fmt"
)

func main() {
	//测试
	iphone14 := model.Electronics{
		model.Goods{"iphone", 12000, 100},
		"apple",
		"phone",
	}
	iphone14.PrintInfo()
	iphone14.PrintMore()
	flag := iphone14.CheckIsOver(90)
	fmt.Println(flag)
	iphone14.Add(10)
	fmt.Println(iphone14.Inventory)
	iphone14.Reduce(20)
	fmt.Println(iphone14.Inventory)
}
