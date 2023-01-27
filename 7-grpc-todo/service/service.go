package main

import (
	"7-grpc-todo/model"
	"7-grpc-todo/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.ListTodo

// var db *sql.DB

func init() {
	localStorage = new(model.ListTodo)
	localStorage.List = make([]*model.Todo, 0)
	// connStr := "postgres://postgresuser:postgrespassword@localhost:5432/postgres?sslmode=verify-full"
	// db, _ = sql.Open("postgres", connStr)
}

type TodosServer struct {
	model.UnimplementedTodosServer
	repository repository.TodoRepository
}

func (s TodosServer) CreateTodo(ctx context.Context, param *model.Todo) (*emptypb.Empty, error) {
	_, err := s.repository.CreateTodo(param)
	return new(emptypb.Empty), err
}
func (s TodosServer) GetTodos(ctx context.Context, param *emptypb.Empty) (*model.ListTodo, error) {
	var todos *model.ListTodo
	// rows, _ := db.Query(`SELECT * FROM todos`)
	// rows.Scan(&todos)
	return todos, nil
}

func (s TodosServer) GetByID(ctx context.Context, param *model.InputTodoID) (*model.Todo, error) {
	data, err := s.repository.GetByID(param.Id)
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println(data)
	return data, nil
}
func (s TodosServer) UpdateTodo(ctx context.Context, param *model.Todo) (*model.Todo, error) {
	data, err := s.repository.UpdateTodo(param.Id, param)
	if err != nil {
		log.Panicln(err.Error())
	}
	return data, nil
}

func (s TodosServer) DeleteTodo(ctx context.Context, param *model.InputTodoID) (*emptypb.Empty, error) {
	err := s.repository.DeleteByID(param.Id)
	return &emptypb.Empty{}, err
}

func main() {
	server := grpc.NewServer()
	// var todoServer TodosServer
	connStr := "postgres://postgresuser:postgrespassword@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(db)
	todoRepository := repository.NewRepository(db)
	todoServer := TodosServer{repository: todoRepository}
	model.RegisterTodosServer(server, todoServer)

	l, _ := net.Listen("tcp", "localhost:7000")

	server.Serve(l)
}
