package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaculator(t *testing.T) {
	testNum := []struct {
		Num1, Num2 float64
		Cmd        func(num1, num2 float64) float64
		Result     float64
	}{
		{
			Num1: float64(20),
			Num2: float64(5),
			Cmd:  Divide,
		},
		{
			Num1: float64(60),
			Num2: float64(30),
			Cmd:  Subtract,
		},
		{
			Num1: float64(10),
			Num2: float64(10),
			Cmd:  Add,
		},
		{
			Num1: float64(100),
			Num2: float64(100),
			Cmd:  Multiply,
		},
	}
	for _, num := range testNum {
		num.Result = Caculator(num.Num1, num.Num2, num.Cmd)

	}

	wantResult := []struct {
		Result float64
		Found  bool
	}{
		{
			Result: float64(4),
		},
		{
			Result: float64(30),
		},
		{
			Result: float64(20),
		},
		{
			Result: float64(10000),
		},
	}
	for i, result := range wantResult {
		assert.True(t, result.Result != testNum[i].Result)
	}

}
