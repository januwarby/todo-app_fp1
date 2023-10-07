package todo_repository

import (
	"todo-app_fp1/entity"
	"todo-app_fp1/pkg/errs"
)

type TodoRepository interface {
	CreateTodo(todo *entity.Todo) (*entity.Todo, errs.MessageErr)
	GetAllTodos() ([]entity.Todo, errs.MessageErr)
	GetTodoByID(id uint) (*entity.Todo, errs.MessageErr)
	UpdateTodo(oldTodo *entity.Todo, newTodo *entity.Todo) (*entity.Todo, errs.MessageErr)
	DeleteTodo(id uint) errs.MessageErr
}
