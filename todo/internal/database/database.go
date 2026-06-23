package database  

import (
	"time" 
	"log" 
	"database/sql"
	_ "modernc.org/sqlite" 
)

type Database struct{
	DB *sql.DB 
}


func NewDB() *sql.DB {
	dbPointer, err := sql.Open("sqlite", "/home/khodges/dev/golang/todolist-cli/todo/list_database.db") 
	if err != nil {
		log.Panic(err)
	}
	return dbPointer   
}

func (d *Database) AddTodo(title string)  {
	stmt, err := d.DB.Prepare("INSERT INTO lists (task, date, completed) VALUES( ?, ?, ?)")
  if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()
	t := time.Now() 
	_, err = stmt.Exec(title, t.Format("2006-01-02 15:04:05"), 0)
	if err != nil {
		log.Panic(err)
	}
}  

func (d *Database) ListTodo() *sql.Rows {
	rows, err := d.DB.Query("SELECT * FROM lists")
	if err != nil {
		log.Panic(err)
	}
	return rows 
}

func (d *Database) DeleteTodo(id int) {
	stmt, err := d.DB.Prepare("DELETE FROM lists WHERE id= ?")
	if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()

	stmt.Exec(id)
}

func (d *Database) CompleteTodo(id int) {
	stmt, err := d.DB.Prepare("UPDATE lists SET completed=1 WHERE id = ?") 
	if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()

	stmt.Exec(id)
}
