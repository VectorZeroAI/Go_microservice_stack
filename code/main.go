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
type Log_entry struct {
	App string
	Action string
	TimeStamp time.Time // The input will increment this.
}
// per app log organisation
type per_app_log struct {
	App string
	Logs []Log_entry
}
// full log storage
var FullLog []per_app_log

// The input data structure
type data struct {
	Application string `json:"application"`
	Action_code int `json:"action_code"`
}

/*
Map of action codes to actions:
1. : start of a programm
2. : init of logger
3. : log an event
4. : non_fatal error
5. : fatal error
6. : shutdown
99. : log deletion
*/

func postDataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON body
	var d data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	// Now We check what the action code is, and act accordingly
	if d.Action_code == 1 {
		tmp := per_app_log{App: d.Application, Logs: []Log_entry{}}
		current_log_entry := Log_entry{
			App: d.Application,
			Action: "App started",
			TimeStamp: time.Now(),
		}
		tmp.Logs = append(tmp.Logs, current_log_entry)
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
