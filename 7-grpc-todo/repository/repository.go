package repository

import (
	"7-grpc-todo/model"
	"database/sql"
	"fmt"
)

type TodoRepository interface {
	CreateTodo(newTodo *model.Todo) (*model.Todo, error)
	GetByID(id int32) (*model.Todo, error)
	UpdateTodo(id int32, updateTodo *model.Todo) (*model.Todo, error)
	DeleteByID(id int32) error
}

type todoRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) CreateTodo(newTodo *model.Todo) (*model.Todo, error) {
	fmt.Println("Masuk nih")
	var todoId int32
	err := tr.db.QueryRow(`
		INSERT INTO todos(name)
		VALUES('testing') RETURNING id`).Scan(&todoId)
	fmt.Println(err)
	if err != nil {
		return newTodo, err
	}
	newTodo.Id = todoId
	return newTodo, nil
}
func (tr *todoRepository) GetByID(id int32) (*model.Todo, error) {
	var foundTodo model.Todo
	rows := tr.db.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)
	rows.Scan(&foundTodo.Id, &foundTodo.Name)
	return &foundTodo, nil
}
func (tr *todoRepository) UpdateTodo(id int32, updateTodo *model.Todo) (*model.Todo, error) {
	_, err := tr.db.Query(`UPDATE todos SET name = $1 WHERE id = $2`, updateTodo.Name, id)
	if err != nil {
		return nil, err
	}
	return updateTodo, nil
}

func (tr *todoRepository) DeleteByID(id int32) error {
	_, err := tr.db.Query(`DELETE from todos WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
