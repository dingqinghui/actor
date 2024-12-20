// Package actor
// @Description: 无锁消息队列

package actor

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	next *node
	val  interface{}
}

type Queue struct {
	head, tail *node
}

func NewQueue() *Queue {
	q := &Queue{}
	stub := &node{}
	q.head = stub
	q.tail = stub
	return q
}

func (q *Queue) Push(x interface{}) {
	n := new(node)
	n.val = x
	prev := (*node)(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), unsafe.Pointer(n)))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&prev.next)), unsafe.Pointer(n))
}

func (q *Queue) Pop() interface{} {
	tail := q.tail
	next := (*node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next)))) // acquire
	if next != nil {
		q.tail = next
		v := next.val
		next.val = nil
		return v
	}
	return nil
}

func (q *Queue) Empty() bool {
	tail := q.tail
	next := (*node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next))))
	return next == nil
}
