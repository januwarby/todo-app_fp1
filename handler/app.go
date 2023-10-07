package handler

import (
	"log"
	"os"
	"todo-app_fp1/database"
	"todo-app_fp1/docs"
	"todo-app_fp1/handler/http_handler"
	"todo-app_fp1/repository/todo_repository/todo_pg"
	"todo-app_fp1/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	port    string
	appHost = os.Getenv("APP_HOST")
	ginMode = os.Getenv("GIN_MODE")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Mengambil nilai PORT dari .env, jika tidak ada, akan menjadi string kosong
	port = os.Getenv("PORT")

	docs.SwaggerInfo.Title = "Todo Application"
	docs.SwaggerInfo.Description = "This is a todo list management application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	if ginMode == "release" {
		docs.SwaggerInfo.Host = appHost
	} else {
		docs.SwaggerInfo.Host = appHost + ":" + port
	}
}

func StartApp() {
	db := database.GetDBInstance()

	r := gin.Default()

	todoRepo := todo_pg.NewTodoPG(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := http_handler.NewTodoHandler(todoService)

	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos", todoHandler.GetAllTodos)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if port == "" {
		log.Fatal("PORT is not defined in the .env file")
	} else {
		log.Fatalln(r.Run(":" + port))
	}
}
