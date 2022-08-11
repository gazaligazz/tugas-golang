package models

import (
	"database/sql"
	"fmt"
	"time"

	"crud-task/config"
	"crud-task/entities"
)

type TaskingModel struct {
	conn *sql.DB
}

func NewTaskingModel() *TaskingModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &TaskingModel{
		conn: conn,
	}
}

func (p *TaskingModel) FindAll() ([]entities.Tasking, error) {

	rows, err := p.conn.Query("select * from tasking")
	if err != nil {
		return []entities.Tasking{}, err
	}
	defer rows.Close()

	var dataTasking []entities.Tasking
	for rows.Next() {
		var tasking entities.Tasking
		rows.Scan(&tasking.Id,
			&tasking.Task,
			&tasking.Assignee,
			&tasking.Deadline)

		tgl_lahir, _ := time.Parse("2006-01-02", tasking.Deadline)
		tasking.Deadline = tgl_lahir.Format("02-01-2006")
		dataTasking = append(dataTasking, tasking)
	}

	return dataTasking, nil

}

func (p *TaskingModel) Create(tasking entities.Tasking) bool {

	result, err := p.conn.Exec("insert into tasking (task, assignee, deadline) values(?,?,?)",
		tasking.Task, tasking.Assignee, tasking.Deadline)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *TaskingModel) Find(id int64, tasking *entities.Tasking) error {

	return p.conn.QueryRow("select * from tasking where id = ?", id).Scan(
		&tasking.Id,
		&tasking.Task,
		&tasking.Assignee,
		&tasking.Deadline)
}

func (p *TaskingModel) Update(tasking entities.Tasking) error {

	_, err := p.conn.Exec(
		"update tasking set task = ?, assignee = ?, deadline = ? where id = ?",
		tasking.Task, tasking.Assignee, tasking.Deadline, tasking.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *TaskingModel) Delete(id int64) {
	p.conn.Exec("delete from tasking where id = ?", id)
}
