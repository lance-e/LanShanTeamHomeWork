package synchronousMap

import (
	"fmt"
	"testing"
	"time"
)

func TestConMap_sync(t *testing.T) {
	value := []any{"first", "second", "third", "four", "five"}
	key := []int{1, 2, 3, 4, 5}
	m1 := NewConMap()
	go func() {

		m1.Put(key[0], value[0])
		m1.Put(key[1], value[1])
		time.Sleep(2 * time.Second)
		m1.Put(key[2], value[2])
	}()
	go func() {
		time.Sleep(time.Second)
		value1, _ := m1.Get(key[1], 5*time.Second)
		fmt.Println(value1)
		//time.Sleep(5 * time.Second)
		if value, err := m1.Get(key[2], 5*time.Second); err != nil {
			fmt.Println(value)

		}
	}()
	time.Sleep(10 * time.Second)
}
