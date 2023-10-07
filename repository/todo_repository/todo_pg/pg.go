package todo_pg

import (
	"fmt"
	"log"
	"todo-app_fp1/entity"
	"todo-app_fp1/pkg/errs"
	"todo-app_fp1/repository/todo_repository"

	"gorm.io/gorm"
)

type todoPG struct {
	db *gorm.DB
}

func NewTodoPG(db *gorm.DB) todo_repository.TodoRepository {
	return &todoPG{db: db}
}

func (t *todoPG) CreateTodo(todo *entity.Todo) (*entity.Todo, errs.MessageErr) {
	if err := t.db.Create(todo).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to create new todo")
	}

	return todo, nil
}

func (t *todoPG) GetAllTodos() ([]entity.Todo, errs.MessageErr) {
	var todos []entity.Todo

	if err := t.db.Find(&todos).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to get all todos")
	}

	return todos, nil
}

func (t *todoPG) GetTodoByID(id uint) (*entity.Todo, errs.MessageErr) {
	var todo entity.Todo
	if err := t.db.First(&todo, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Todo with id %v is not found", id))
	}
	return &todo, nil
}

func (t *todoPG) UpdateTodo(oldTodo *entity.Todo, newTodo *entity.Todo) (*entity.Todo, errs.MessageErr) {
    oldTodo.Title = newTodo.Title
    
    // Menggunakan nilai newTodo.Completed jika true, jika tidak, maka set ke false
    if newTodo.Completed {
        oldTodo.Completed = true
    } else {
        oldTodo.Completed = false
    }

    if err := t.db.Save(oldTodo).Error; err != nil {
        return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update todo with id %v", oldTodo.ID))
    }
    return oldTodo, nil
}



func (t *todoPG) DeleteTodo(id uint) errs.MessageErr {
	if err := t.db.Where("id = ?", id).Delete(&entity.Todo{}).Error; err != nil {
		log.Println(err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete todo with id %v", id))
	}
	return nil
}
