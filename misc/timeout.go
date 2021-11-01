package misc

import "time"

type Timeout struct {
	Released bool
	Occured  bool
}

func (t *Timeout) Run(seconds int, callback func()) {
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		if !t.Released {
			t.Occured = true
			callback()
		}
	}()
}

func (t *Timeout) Release() {
	t.Released = true
}

// generate
