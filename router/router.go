package router

import ( 
	"github.com/gin-gonic/gin"
	"example.com/task_manager/controllers"
)



func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/tasks", controllers.CreateTask)
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/tasks/:id", controllers.GetTaskByID)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)	
	}

	return r
}