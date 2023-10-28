package utile

import (
	"fmt"
	"math"
)

// LV2
// 编写一个Go函数，接受圆的半径作为参数，然后返回圆的面积。使用 math 包中的常数 Pi。在 main 函数中调用此函数并打印结果。
//
// 提示，引入 Pi 只需要写出math.Pi
func Area(r float64) {
	fmt.Println(math.Pi * r * r)
}
