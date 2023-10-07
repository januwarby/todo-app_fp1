package http_handler

import (
	"net/http"
	"strconv"
	"todo-app_fp1/dto"
	"todo-app_fp1/pkg/errs"
	"todo-app_fp1/service"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *todoHandler {
	return &todoHandler{todoService: todoService}
}

// CreateTodo godoc
//
//	@Summary		Create a todo
//	@Description	Create a todo by json
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		dto.NewTodoRequest	true	"Create todo request body"
//	@Success		201		{object}	dto.NewTodoResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/todos [post]
func (t *todoHandler) CreateTodo(ctx *gin.Context) {
	var requestBody dto.NewTodoRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdTodo, err := t.todoService.CreateTodo(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdTodo)
}

// GetAllTodos godoc
//
//	@Summary		Get all todos
//	@Description	List todos
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.GetAllTodosResponse
//	@Failure		500	{object}	errs.MessageErrData
//	@Router			/todos [get]
func (t *todoHandler) GetAllTodos(ctx *gin.Context) {
	todos, err := t.todoService.GetAllTodos()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
//
//	@Summary		Get a todo
//	@Description	Get a todo by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Todo ID"
//	@Success		200	{object}	dto.GetTodoByIDResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Router			/todos/{id} [get]
func (t *todoHandler) GetTodoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newError := errs.NewBadRequest("ID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	todo, err2 := t.todoService.GetTodoByID(uint(idUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
//
//	@Summary		Update todo
//	@Description	Update a todo by json
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint				true	"Todo ID"
//	@Param			todo	body		dto.NewTodoRequest	true	"Update todo request body"
//	@Success		200		{object}	dto.GetTodoByIDResponse
//	@Failure		400		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		404		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/todos/{id} [put]
func (t *todoHandler) UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newError := errs.NewBadRequest("ID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var newTodo dto.NewTodoRequest

    if err := ctx.ShouldBindJSON(&newTodo); err != nil {
        newError := errs.NewUnprocessableEntity(err.Error())
        ctx.JSON(newError.StatusCode(), newError)
        return
    }

	updatedTodo, err2 := t.todoService.UpdateTodo(uint(idUint), &newTodo)
    if err2 != nil {
        ctx.JSON(err2.StatusCode(), err2)
        return
    }

	ctx.JSON(http.StatusOK, updatedTodo)
}

// DeleteTodo godoc
//
//	@Summary		Delete todo
//	@Description	Delete a todo by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"todo ID"
//	@Success		200	{object}	dto.DeleteTodoResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Failure		500	{object}	errs.MessageErrData
//	@Router			/todos/{id} [delete]
func (t *todoHandler) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newError := errs.NewBadRequest("ID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err2 := t.todoService.DeleteTodo(uint(idUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
