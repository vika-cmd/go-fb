package note

import "time"

type ModelNoteForm struct {
	Task        string
	Category    string
	Priority    int
	Description string
	ByDate      time.Time
}