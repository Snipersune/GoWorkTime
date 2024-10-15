package timeentry

import "time"

type TimeEntry struct {
	Date     time.Time `json:"date"`
	Duration float64   `json:"duration"`
	Task     string    `json:"task"`
}

func New(date time.Time, duration float64, task string) TimeEntry {
	return TimeEntry{
		Date:     date,
		Duration: duration,
		Task:     task,
	}
}
