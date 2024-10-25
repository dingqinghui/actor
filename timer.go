/**
 * @Author: dingQingHui
 * @Description:
 * @File: timer
 * @Version: 1.0.0
 * @Date: 2024/10/24 18:20
 */

package actor

import (
	"github.com/RussellLuo/timingwheel"
	"time"
)

var (
	tw *timingwheel.TimingWheel
)

func init() {
	tw = timingwheel.NewTimingWheel(10*time.Millisecond, 3600)
	tw.Start()
}

type twTimer struct {
	t   *timingwheel.Timer
	id  int32
	fn  func()
	hub *timerHub
}

func (t *twTimer) Id() int32 {
	return t.id
}
func (t *twTimer) Trigger() {
	if t.fn != nil {
		t.fn()
	}
}
func (t *twTimer) Stop() {
	t.hub.Remove(t.Id())
	t.t.Stop()
}

func newTimerHub(process IProcess) ITimerHub {
	return &timerHub{
		timerDict: make(map[int32]ITimer),
		process:   process,
	}
}

type timerHub struct {
	timerDict map[int32]ITimer
	process   IProcess
	id        int32
}

func (th *timerHub) nextId() int32 {
	th.id++
	return th.id
}
func (th *timerHub) AddTimer(d time.Duration, fn func()) ITimer {
	id := th.nextId()
	t := tw.AfterFunc(d, func() {
		_ = th.process.Send(&TimerMessage{id: id})
	})
	twt := &twTimer{
		fn: fn,
		id: id,
		t:  t,
	}
	th.timerDict[id] = twt
	return twt
}

func (th *timerHub) Get(id int32) ITimer {
	twt, ok := th.timerDict[id]
	if !ok {
		return nil
	}
	return twt
}

func (th *timerHub) Remove(id int32) {
	delete(th.timerDict, id)
}
