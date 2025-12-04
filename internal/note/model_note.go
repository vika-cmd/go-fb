package note

import "time"

type ModelNoteForm struct {
	Task        string
	Category    string
	Priority    int
	Description string
	ByDate      time.Time
	Createdat  	time.Time
}

/* type ModelNoteDb struct {
	Id
	Task        string
	Category    string
	Priority    int
	Description string
	ByDate      time.Time
	Createdat  	time.Time
	Readydat 	time.Time
} */

