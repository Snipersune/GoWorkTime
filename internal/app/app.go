package app

import (
	"os"
)

func Run() {
	if len(os.Args) < 1 {
		return
	}
	var args = os.Args[1:]

	switch args[0] {
	case "start":
		start(args)
	case "stop":
		stop()
	case "log":
		log()
	case "status":
		status()
	case "report":
		report()
	default:
		help()
	}

}

func start(args []string) {
	var appState, err = LoadAppState()
	if err != nil {
		return
	}

	if !appState.IsValidActiveLog() {
		println("No active log to start. Use 'gwt log --list' to see available logs.")
		return
	}
	if appState.IsRecording {
		println("Active log is already being recorded to.")
		return
	}

	// Open log file
	logFile, err := storageOpenLog(appState.ActiveLog)
	if err != nil {
		println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	appState.IsRecording = true
	appState.Save()
}

// gwt start
// gwt stop
// gwt stop --message <message>

// gwt report
// gwt report -a (default)
// gwt report --all (default)
// gwt report --active
// gwt report --session <ID>

// gwt session -m <message>
// gwt session --message <message>
// gwt session --list
// gwt session --add <ID>
// gwt session --remove <ID>
// gwt session --modify <ID>

// gwt log <name>
// gwt log -l
// gwt log --list
// gwt log -r <name>
// gwt log --rm <name>

// gwt switch <name>
// gwt switch -c <name>
// gwt switch --create <name>

// gwt status
// gwt help
