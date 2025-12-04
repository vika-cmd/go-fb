package note

import (
	//"context"
	"fmt"
	//"time"

	//"github.com/jackc/pgx/v5"
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
fmt.Printf("repo#addNote- Task - %s", form.Task)
/* 	query := `INSERT INTO tasks (task, email, role, category, priority, description, readydat) VALUES (@task, @email, @role, @category, @priority, @description, @readydat)`
	args := pgx.NamedArgs{
		"task": form.Task,
		"email": form.Email,
		"role": form.Role,
		"category": form.Category,
		"priority": form.Priority,
		"description": form.Description,
		"readydat": form.Readydat,
		"createdat": time.Now(),		
	}
	_,err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Printf("repo#addTodo- err - %s", err)		
	} */
	
	//fmt.Printf("repo#addTodo: Тип  form.Readydata: %T\n", form.Readydat)  time.Time  time.Now()

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

