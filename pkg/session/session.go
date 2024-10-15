package session

import (
	"github.com/Snipersune/GoWorkTime/pkg/timeentry"
)

type Session struct {
	Log       string
	TimeEntry timeentry.TimeEntry
}

func (session Session) GetTimeEntry() timeentry.TimeEntry {
	return session.TimeEntry
}

func (session Session) GetLog() string {
	return session.Log
}

func (session Session) SetLog(log string) {
	session.Log = log
}
