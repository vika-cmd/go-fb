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

type ModelNoteDb struct {
	Id 			int		 `db:"id"`
	Task        string	 `db:"task"`
	Category    string 	 `db:"category"`
	Priority    int 	 `db:"priority"`
	Description string 	 `db:"description"`
	ByDate      time.Time	`db:"bydate"`
	Createdat  	time.Time 	`db:"createdat"`
	Readydat 	time.Time 	`db:"readydat"`
}

