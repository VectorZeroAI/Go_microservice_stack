package main
// Later to be revised into package loger, and a loader system to be used.
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// --------- Types creation  -------------
// A singe log entry
type LogEntry struct {
	App string
	Action string
	TimeStamp time.Time // The input will increment this.
}
// per app log organisation
type per_app_log struct {
	App string
	Logs []LogEntry
	IsShutdown bool
}
// full log storage
var FullLog []per_app_log

// The input IncommingString structure
type IncommingString struct {
	Application string `json:"application"`
	Action_code int `json:"action_code"`
	Event string `json:"Event"`
}

/*
Map of action codes to actions:
1. : start of a programm
3. : log an event
4. : non_fatal error
5. : fatal error
6. : shutdown
99. : log deletion
*/


func postDataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON body
	var CurrentIncommingString IncommingString
	err := json.NewDecoder(r.Body).Decode(&CurrentIncommingString);
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	// Now We check what the action code is, and act accordingly
	if CurrentIncommingString.Action_code == 1 {
		tmp := per_app_log{App: CurrentIncommingString.Application, Logs: []LogEntry{}, IsShutdown: false}
		current_log_entry := LogEntry{
			App: CurrentIncommingString.Application,
			Action: "App started",
			TimeStamp: time.Now(),
		}
		tmp.Logs = append(tmp.Logs, current_log_entry)
		FullLog = append(FullLog, tmp)
		return
	}
	Log_string := func () {
			tmp := LogEntry{
				App: CurrentIncommingString.Application,
				Action: CurrentIncommingString.Event,
				TimeStamp: time.Now(),
			}
			for i := 0; i < len(FullLog); i++ {
				if FullLog[i].App == CurrentIncommingString.Application {
					FullLog[i].Logs = append(FullLog[i].Logs, tmp)
					return
				}
			}
		}
	if CurrentIncommingString.Action_code == 3 {Log_string()}
	if CurrentIncommingString.Action_code == 4 {Log_string()}
	if CurrentIncommingString.Action_code == 5 {
		tmp := LogEntry{
			App: CurrentIncommingString.Application,
			Action: CurrentIncommingString.Event,
			TimeStamp: time.Now(),
		}
		for i := 0; i < len(FullLog); i++ {
			if FullLog[i].App == CurrentIncommingString.Application {
				FullLog[i].Logs = append(FullLog[i].Logs, tmp)
				FullLog[i].IsShutdown = true
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	Err := http.ListenAndServe(":8080", nil)
	if Err != nil {
		fmt.Println("Error starting server:", Err)
	}
}
