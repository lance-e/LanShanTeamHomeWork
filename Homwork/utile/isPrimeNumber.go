package utile

import "fmt"

// LV3
// 编写一个Go函数，接受一个整数作为参数，然后判断它是否为素数（质数）。在 main 函数中调用此函数并打印结果。提示：一个素数是只能被 1 和自身整除的正整数。
func IsPrimeNumber(n int) {
	if n <= 0 {
		fmt.Println("负数绝壁不可能是素数")
		return
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println(n, " 不是素数")
			return
		}
	}
	fmt.Println(n, " 是素数")
}
