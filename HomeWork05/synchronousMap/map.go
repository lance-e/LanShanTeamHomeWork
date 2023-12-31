package synchronousMap

import (
	"errors"
	"fmt"
	"time"
)

//var a sync.Map

// conMap 既然是要求查询时若 key 不存在则阻塞并等待 maxWaitingTime 时间，
// 若在该时间内有写入该 key 的操作，返回 value，超时则返回超时错误
// 所以必然不是并发安全的
type conMap struct {
	m   map[int]any
	ch1 chan int
	ch2 chan any
}

func NewConMap() *conMap {
	return &conMap{
		m:   make(map[int]any),
		ch1: make(chan int),
		ch2: make(chan any),
	}
}

type syncMap interface {
	Get(key int, time time.Duration) (any, error)
	Put(key int, value any)
}

func (m *conMap) Get(key int, maxWaitingTime time.Duration) (any, error) {
	value, ok := m.m[key]
	if !ok {
		now := time.Now()
		for {
			if time.Since(now) >= maxWaitingTime {
				return nil, errors.New("超时")
			}
			select {
			case <-m.ch1:
				go func() {
					data := <-m.ch2
					m.m[key] = data
				}()
				return m.m[key], nil
			default:
				fmt.Println("waiting for enter value")
				time.Sleep(time.Second)
			}
		}
	}

	return value, nil
}
func (m *conMap) Put(key int, value any) bool {
	if _, ok := m.m[key]; ok {
		return false
	}
	m.m[key] = value
	m.ch1 <- key

	return true
}
