package peperson

type Mutax interface {
	Lock()
	UnLock()
}
type Locker struct {
	wantThread [2]bool
	turn       int
	self       int
}

func NewLocker() *Locker {
	return &Locker{
		wantThread: [2]bool{false, false},
		turn:       0,
	}
}

func (l *Locker) Lock() {
	l.wantThread[l.self] = true //设置自己想进入临界区
	l.turn = 1 - l.self         //
	//如果对面也想进入临界区，并且此时轮到对面了，就死循环让自己等待
	for l.wantThread[1-l.self] && l.turn == 1-l.self {
	}
}
func (l *Locker) UnLock() {
	l.wantThread[l.self] = false //表示自己退出
}
