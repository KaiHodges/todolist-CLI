package internal 

import (
	"database/sql"
	_ "modernc.org/sqlite" 
)

type database struct{
	db *sql.DB 
}

func (d *database) DbConnect() error {
	dbPointer, err := sql.Open("sqlite", "list_database.db") 
	d.db = dbPointer 
	return err 
}

func (d *database) AddTodo(title string) error {


}  

func (d *database) ListTodo() error {

}

func (d *database) DeleteTodo(id int) error {

}

func (d *database) CompleteTodo(id int) error {

}
