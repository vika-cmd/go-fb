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
fmt.Println("repo#addNote-", form.Task, form.Category, form.Priority, form.Description, form.ByDate, time.Now())
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

fmt.Printf("repo#addNote: Тип  form.ByDate: %T\n", form.ByDate)
fmt.Println("repo#addNote: time.Now(): ", time.Now())

return nil
} 



/* func (r *RepositoryNote) CountAll() int{
	query := "SELECT count(*) FROM notebook"
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
} */

/* func (r *RepositoryNote) GetAll(limit, offset int) ([]TodoDB, error ){
	query := `SELECT * FROM tasks ORDER BY createdat LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit": limit,
		"offset": offset,
	}

	rows, err := r.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}

	// ver new
	todoes, err := pgx.CollectRows(rows, pgx.RowToStructByName[TodoDB])
	if err != nil {
		return nil, err
	}
	return todoes,nil */

	// ver old
/* 	for rows.Next(){
	    var row TodoDB
		rows.Scan(row)
	} */
	 
//}

