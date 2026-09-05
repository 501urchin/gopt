package gopt

type Semaphore struct {
	sema chan struct{}
}

func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		sema: make(chan struct{}, limit),
	}
}

func (s *Semaphore) Release() {
	<-s.sema
}
func (s *Semaphore) Acquire() {
	s.sema <- struct{}{}
}