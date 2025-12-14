package note

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryNote struct {
	Dbpool *pgxpool.Pool
}

func NewRepositoryNote(dbpool *pgxpool.Pool) *RepositoryNote{
	return &RepositoryNote{
		Dbpool: dbpool,
	}
}



func (r *RepositoryNote) addNote(form ModelNoteForm) error{
	/* ---- */
	query := `INSERT INTO note (task, category, priority, description, bydate,createdat) VALUES (@task, @category, @priority, @description, @bydate, @createdat)`
	args := pgx.NamedArgs{
		"task": form.Task,
		"category": form.Category,
		"priority": form.Priority,
		"description": form.Description,
		"bydate": form.ByDate,		
		"createdat": time.Now(),		
	}
	_,err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Printf("repo#addNote- err - %v\n", err)
		return err		
	}
	/* ---- */

	//fmt.Printf("repo#addNote: Тип  form.ByDate: %T\n", form.ByDate)
	//fmt.Println("repo#addNote: time.Now(): ", time.Now())

	return nil
} 

// COALESCE(date_column, '0001-01-01'::date)
//	query := `SELECT id, task, category, priority, description, bydate, createdat, COALESCE(readydat, '0001-01-01'::date) FROM note`
//LIMIT 20 OFFSET 40 вернет 20 записей, начиная с 41-й
func (r *RepositoryNote) GetAll(limit, offset int) ([]ModelNoteDb,error) {
	query := `SELECT * FROM note ORDER BY bydate LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit": limit,
		"offset": offset,
	}
	rows, err := r.Dbpool.Query(context.Background(), query, args)
	
	if err != nil {
		fmt.Printf("repo#GetAll: query - err: %v\n",err)
		return nil, err
	}
	//---------- ?
	defer rows.Close()

	// ver new   cannot scan null into *time.Time 
	notes, err := pgx.CollectRows(rows, pgx.RowToStructByName[ModelNoteDb])
	if err != nil {
		fmt.Printf("repo#GetAll: CollectRows - err: %v\n",err)
		return nil, err
	}
	return notes,nil 

	// ver OLD ?
	//var arrNotes []ModelNoteDb
	//var row ModelNoteDb
/* 	for rows.Next(){
		//var row ModelNoteDb number of field descriptions must equal number of  destination, got 7 and 1
		var task string
		err := rows.Scan(&task)
		if err !=nil {
			fmt.Printf("repo#GetAll: Scan - err: %v\n",err)
		}
		fmt.Println("repo#GetAll: rows.Next - row: ",row)
		fmt.Printf("repo#GetAll: rows.Next - rows: %v\n",rows)
		arrNotes = append(arrNotes, row)
		fmt.Printf("repo#GetAll: rows.Next - arrNotes: %v\n",arrNotes)
	
	} */
	//return arrNotes, err
}

func (r *RepositoryNote) TotalRecords() int{
	query := "SELECT count(*) FROM note"
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}