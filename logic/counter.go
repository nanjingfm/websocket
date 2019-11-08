package logic

import "sync/atomic"

var _defaultCounter *Counter

func InitCounter() {
	_defaultCounter = &Counter{}
}

func GetDefaultCounter() *Counter {
	return _defaultCounter
}

type Counter struct {
	MsgCount uint64 // 消息数量
	MsgSize  uint64 // 消息体大小
}

func (c *Counter) IncreaseMsgCount(n uint64) {
	_ = atomic.AddUint64(&(c.MsgCount), n)
}

func (c *Counter) IncreaseMsgSize(size uint64) {
	_ = atomic.AddUint64(&(c.MsgSize), size)
}
