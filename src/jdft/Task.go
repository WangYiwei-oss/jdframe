package jdft

import (
	"sync"
)

type TaskFunc func(params ...interface{})

var jdTaskChan chan *JdTask
var once sync.Once

func init() {
	go func() {
		for {
			jt := <-getJdTaskChan()
			jt.Exec()
		}
	}()
}

func getJdTaskChan() chan *JdTask {
	once.Do(func() {
		jdTaskChan = make(chan *JdTask)
	})
	return jdTaskChan
}

type JdTask struct {
	f  TaskFunc
	ps []interface{}
	cb func()
}

func NewJdTask(f TaskFunc, ps []interface{}, cb func()) *JdTask {
	return &JdTask{f: f, ps: ps, cb: cb}
}

func (j *JdTask) Exec() {
	go func() {
		defer j.cb()
		j.f(j.ps...)
	}()
}

func DoTask(f TaskFunc, cb func(), ps ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getJdTaskChan() <- NewJdTask(f, ps, cb)
	}()
}
