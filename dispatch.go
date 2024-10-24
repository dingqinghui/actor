// Package actor
// @Description:

package actor

// 协程调度器
type goroutineDispatcher int

func NewDefaultDispatcher(throughput int) IDispatcher {
	return goroutineDispatcher(throughput)
}
func (goroutineDispatcher) Schedule(fn func(), recoverFun func(err interface{})) error {
	return Submit(fn, recoverFun)
}

func (d goroutineDispatcher) Throughput() int {
	return int(d)
}

// 同步调度器
type synchronizedDispatcher int

func (synchronizedDispatcher) Schedule(fn func(), recoverFun func(err interface{})) error {
	Try(fn, recoverFun)
	return nil
}

func (d synchronizedDispatcher) Throughput() int {
	return int(d)
}

func NewSynchronizedDispatcher(throughput int) IDispatcher {
	return synchronizedDispatcher(throughput)
}
