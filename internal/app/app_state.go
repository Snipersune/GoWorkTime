package app

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Snipersune/GoWorkTime/internal/config"
)

type AppState struct {
	Logs        []string
	ActiveLog   string
	IsRecording bool
}

func LoadAppState() (AppState, error) {
	var fAppState, err = os.ReadFile(config.AppStateFilePath)
	if err == nil {
		return AppState{}, err
	}
	var appState = AppState{}
	err = json.Unmarshal(fAppState, &appState)
	if err == nil {
		return AppState{}, err
	}
	return appState, err
}

func (appState AppState) AddLog(log string) error {
	if ok, _ := appState.InLogs(log); ok {
		return errors.New("log already exists")
	}
	appState.Logs = append(appState.Logs, log)
	return nil
}

func (appState AppState) RemoveLog(log string) error {
	if ok, idx := appState.InLogs(log); ok {
		if appState.ActiveLog == log {
			appState.ActiveLog = ""
			appState.IsRecording = false
		}
		appState.Logs = append(appState.Logs[:idx], appState.Logs[idx+1:]...)
		return nil
	}
	return errors.New("log not found")
}

func (appState AppState) GetLogs() []string {
	return appState.Logs
}

func (appState AppState) InLogs(log string) (bool, int) {
	for i, l := range appState.Logs {
		if l == log {
			return true, i
		}
	}
	return false, 0
}

func (appState AppState) SetActiveLog(log string) error {
	if appState.ActiveLog == log {
		return nil
	}

	if ok, _ := appState.InLogs(log); !ok {
		appState.Logs = append(appState.Logs, log)
	}
	appState.ActiveLog = log

	return appState.Save()
}

func (appState AppState) GetActiveLog() string {
	return appState.ActiveLog
}

func (appState AppState) IsValidActiveLog() bool {
	return appState.ActiveLog != ""
}

func (appState AppState) Save() error {
	data, err := json.Marshal(appState)
	if err != nil {
		return err
	}
	return os.WriteFile(config.AppStateFilePath, data, 0644)
}
