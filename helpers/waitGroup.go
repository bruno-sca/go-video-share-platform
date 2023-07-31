package helpers

import "sync"

type WaitGroupHelper struct {
	sync.WaitGroup
}

func (w *WaitGroupHelper) Wrap(f func()) {
	w.Add(1)
	go func() {
		f()
		w.Done()
	}()
}