package dto

import (
	"time"
	"todo-app_fp1/entity"
)

type NewTodoRequest struct {
	Title     string `json:"title" binding:"required" example:"Belajar Golang"`
	Completed bool   `json:"completed" example:"false"`
}

func (t *NewTodoRequest) TodoRequestToEntity() *entity.Todo {
	return &entity.Todo{
		Title:     t.Title,
		Completed: t.Completed,
	}
}

type NewTodoResponse struct {
	Message string         `json:"message" example:"Todo with id 69 has been successfully created"`
	Data    NewTodoRequest `json:"data"`
}

type GetAllTodosResponse struct {
	Message string     `json:"message" example:"success"`
	Data    []TodoData `json:"data"`
}

type GetTodoByIDResponse struct {
	Message string           `json:"message" example:"success"`
	Data    TodoDataDetailed `json:"data"`
}

type DeleteTodoResponse struct {
	Message string `json:"message" example:"Todo with id 5 has been successfully deleted"`
}

type TodoData struct {
	ID        uint   `json:"id" example:"69"`
	Title     string `json:"title" example:"Ngoding"`
	Completed bool   `json:"completed" example:"false"`
}

type TodoDataDetailed struct {
	ID        uint   `json:"id" example:"69"`
	Title     string `json:"title" example:"Ngoding"`
	Completed bool   `json:"completed" example:"false"`
	CreatedAt time.Time `json:"createdAt" example:"2023-04-06T17:55:34.070213+07:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-04-06T17:55:34.070213+07:00"`
}
