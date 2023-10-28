package utile

import (
	"fmt"
	"math"
)

// LVX
// 编写一个Go函数，使用rand包随机选择一个1-100的数（必须每次执行的随机数都不一样），然后使用二分法找到这个数。
func FindNumber(n int64) {
	i := float64(n)
	var middle float64
	low := float64(0)
	high := float64(100)
	for low <= high {
		middle = (high-low)/2 + low
		if i < middle+0.125 && i > middle-0.125 {
			answer := math.Floor(middle + 0.5)
			fmt.Println("计算得出的结果是：", answer)
			break
		} else if i < middle && i >= low {
			high = middle

		} else {
			low = middle

		}
	}
}
