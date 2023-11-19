package routes

import (
	"Memorandum/api"
	_ "Memorandum/docs"
	"Memorandum/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

//不能用gin

func NewRouter() *server.Hertz {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	store := cookie.NewStore([]byte("something-very-secret"))
	url := swagger.URL("http://localhost:8080/swagger/doc.json")
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	h.Use(sessions.New("mysession", store))
	userGroup := h.Group("/api/v1")
	{
		userGroup.POST("user/register", api.Register)
		userGroup.POST("user/login", api.Login)
		task := userGroup.Group("/")
		task.Use(middleware.JWT)
		{
			task.POST("task", api.CreateTask)
			task.PUT("task", api.UpdateTask)
			task.DELETE("task/:id", api.DeleteTask)
			task.DELETE("tasks/:status", api.DeleteTasksInStatus)
			task.DELETE("tasks", api.DeleteAllTasks)
			task.GET("tasks/:status", api.ShowTasksInStatus)
			task.GET("tasks", api.ShowTasksByKey)
		}
	}
	return h
}
