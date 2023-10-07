// service/todo_service.go
package service

import (
	"fmt"
	"todo-app_fp1/dto"
	"todo-app_fp1/pkg/errs"
	"todo-app_fp1/repository/todo_repository"
)

type TodoService interface {
    CreateTodo(payload *dto.NewTodoRequest) (*dto.NewTodoResponse, errs.MessageErr)
    GetAllTodos() (*dto.GetAllTodosResponse, errs.MessageErr)
    GetTodoByID(id uint) (*dto.GetTodoByIDResponse, errs.MessageErr)
    UpdateTodo(id uint, newTodo *dto.NewTodoRequest) (*dto.GetTodoByIDResponse, errs.MessageErr)
    DeleteTodo(id uint) (*dto.DeleteTodoResponse, errs.MessageErr)
}

type todoService struct {
    todoRepo todo_repository.TodoRepository
}

func NewTodoService(todoRepo todo_repository.TodoRepository) TodoService {
    return &todoService{todoRepo: todoRepo}
}

func (s *todoService) CreateTodo(payload *dto.NewTodoRequest) (*dto.NewTodoResponse, errs.MessageErr) {
    todo := payload.TodoRequestToEntity()

    createdTodo, err := s.todoRepo.CreateTodo(todo)
    if err != nil {
        return nil, err
    }

    response := &dto.NewTodoResponse{
        Message: fmt.Sprintf("Todo with id %v has been created successfully", createdTodo.ID),
        Data: dto.NewTodoRequest{
            Title:     createdTodo.Title,
            Completed: createdTodo.Completed,
        },
    }

    return response, nil
}

func (s *todoService) GetAllTodos() (*dto.GetAllTodosResponse, errs.MessageErr) {
    todos, err := s.todoRepo.GetAllTodos()
    if err != nil {
        return nil, err
    }

    todoData := []dto.TodoData{}
    for _, todo := range todos {
        todoData = append(todoData, dto.TodoData{
            ID:        todo.ID,
            Title:     todo.Title,
            Completed: todo.Completed,
        })
    }

    response := &dto.GetAllTodosResponse{
        Message: "success",
        Data:    todoData,
    }

    return response, nil
}

func (s *todoService) GetTodoByID(id uint) (*dto.GetTodoByIDResponse, errs.MessageErr) {
    todo, err := s.todoRepo.GetTodoByID(id)
    if err != nil {
        return nil, err
    }

    response := &dto.GetTodoByIDResponse{
        Message: "success",
        Data: dto.TodoDataDetailed{
            ID:        todo.ID,
            Title:     todo.Title,
            Completed: todo.Completed,
            CreatedAt: todo.CreatedAt,
            UpdatedAt: todo.UpdatedAt,
        },
    }

    return response, nil
}

func (s *todoService) UpdateTodo(id uint, newTodo *dto.NewTodoRequest) (*dto.GetTodoByIDResponse, errs.MessageErr) {
    newTodoEntity := newTodo.TodoRequestToEntity()

    oldTodo, err := s.todoRepo.GetTodoByID(id)
    if err != nil {
        return nil, err
    }

    updatedTodo, err := s.todoRepo.UpdateTodo(oldTodo, newTodoEntity)
    if err != nil {
        return nil, err
    }

    response := &dto.GetTodoByIDResponse{
        Message: fmt.Sprintf("Todo with id %v has been successfully updated", id),
        Data: dto.TodoDataDetailed{
            ID:        updatedTodo.ID,
            Title:     updatedTodo.Title,
            Completed: updatedTodo.Completed,
            CreatedAt: updatedTodo.CreatedAt,
            UpdatedAt: updatedTodo.UpdatedAt,
        },
    }

    return response, nil
}

func (s *todoService) DeleteTodo(id uint) (*dto.DeleteTodoResponse, errs.MessageErr) {
    _, err := s.todoRepo.GetTodoByID(id)
    if err != nil {
        return nil, err
    }

    if err := s.todoRepo.DeleteTodo(id); err != nil {
        return nil, err
    }

    response := &dto.DeleteTodoResponse{
        Message: fmt.Sprintf("Todo with id %v has been successfully deleted", id),
    }

    return response, nil
}
