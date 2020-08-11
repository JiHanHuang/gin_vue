package download

import "sync"

const (
	Waiting int = iota
	Start
	Running
	Finish
	Failed
)

type Attr struct {
	ID       int
	Addr     string
	DownPath string
	FileName string
	FileSize int64
}

type Status struct {
	m       sync.RWMutex
	Percent int
	State   int
}

func InitStatus(s int, p int) *Status {
	return &Status{Percent: p, State: s}
}

func (s *Status) UpS(st int) {
	s.m.Lock()
	s.State = st
	s.m.Unlock()
}

func (s *Status) UpP(p int) {
	s.m.Lock()
	s.Percent = p
	s.m.Unlock()
}

func (s *Status) Read() (rs Status) {
	s.m.RLock()
	rs.Percent = s.Percent
	rs.State = s.State
	s.m.RUnlock()
	return
}

func (s *Status) Write(st int, p int) {
	s.m.Lock()
	s.State = st
	s.Percent = p
	s.m.Unlock()
}
