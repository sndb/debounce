package debounce

import (
	"sync"
	"time"
)

type debouncer struct {
	mu sync.Mutex
	after time.Duration
	timer *time.Timer
}

func New(after time.Duration) func(func()) {
	d := &debouncer{after: after}

	return func(f func()) {
		d.mu.Lock()
		defer d.mu.Unlock()

		if d.timer != nil {
			d.timer.Stop()
		}
		d.timer = time.AfterFunc(d.after, f)
	}
}
