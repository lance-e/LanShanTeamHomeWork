package calculate

func Caculator(num1, num2 float64,cmd func(n1, n2 float64) float64) float64 {
	return cmd(num1,num2)
}

func Add(num1, num2 float64) float64 {
	return num1 + num2
}
func Subtract(num1, num2 float64) float64 {
	return num1 - num2
}
func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}
func Divide(num1, num2 float64) float64 {
	return num1 / num2
}
