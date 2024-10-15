package storage

import (
	"encoding/json"
	"os"

	"github.com/Snipersune/GoWorkTime/internal/config"
	"github.com/Snipersune/GoWorkTime/pkg/timeentry"
)

func SaveTimeEntry(entry timeentry.TimeEntry, filepath string) error {
	var entries []timeentry.TimeEntry

	// Check if file exists, read existing entries
	if _, err := os.Stat(filepath); err == nil {
		data, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		json.Unmarshal(data, &entries)
	}

	// Append new entry
	entries = append(entries, entry)

	json.Marshal()
	// Write back to file
	_, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
}

func LoadTimeEntries(log string) ([]timeentry.TimeEntry, error) {
	var logFname = os.Getenv("HOME") + "/" + config.DataDir + log + ".log"

	var entries []timeentry.TimeEntry
	// Check if file exists, read existing entries
	if _, err := os.Stat(logFname); err == nil {
		data, err := os.ReadFile(logFname)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, &entries)
	}

	// Open log file
	file, err := os.OpenFile(logFname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return file, err
}
