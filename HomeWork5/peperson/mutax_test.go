package peperson

import (
	"fmt"
	"testing"
	"time"
)

func TestLocker_Locker(t *testing.T) {
	locker1 := NewLocker()
	locker2 := NewLocker()
	locker2.self = 1
	locker1.self = 0
	go two()
	locker1.Lock()
	go one()
	locker1.UnLock()
	time.Sleep(time.Second)
}

func one() {
	for i := 0; i < 10; i++ {
		fmt.Println("这是一进程，此时值为：", i)
	}
}
func two() {
	for i := 0; i < 10; i++ {
		fmt.Println("这是二进程，此时值为：", i)
	}
}
