package util

import "sync"

type promise[T any] struct {
	result chan T
	err    chan error
	fin    chan bool
	wg     *sync.WaitGroup
}

func (p promise[T]) Then(callback func(T)) promise[T] {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		select {
		case r := <-p.result:
			callback(r)
		case <-p.fin:
		}
	}()
	return p
}

func (p promise[T]) Catch(callback func(error)) promise[T] {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		select {
		case e := <-p.err:
			callback(e)
		case <-p.fin:
		}
	}()
	return p
}

func (p promise[T]) Wait() {
	p.wg.Wait()
}

func NewPromise[R any](callback func() (R, error)) promise[R] {
	var prom = promise[R]{
		result: make(chan R, 1),
		err:    make(chan error, 1),
		fin:    make(chan bool, 1),
	}
	prom.wg = new(sync.WaitGroup)
	prom.wg.Add(1)
	go func() {
		defer prom.wg.Done()
		result, err := callback()
		if err != nil {
			prom.err <- err
			prom.fin <- false
			return
		}
		prom.result <- result
		prom.fin <- true
	}()
	return prom
}
