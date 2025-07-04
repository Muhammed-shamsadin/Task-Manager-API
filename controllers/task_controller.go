package controllers

import (
	"net/http"

	"example.com/task_manager/data"
	"example.com/task_manager/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	id, err := data.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})

}


func GetTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}


func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := data.GetTaskByID(objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}


func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updateTask := models.UpdateTask{
		// Map fields from task to updateTask as needed, for example:
		Title:       &task.Title,
		Description: &task.Description,
		Status:   &task.Status,
		// Add other fields as required by your UpdateTask struct
	}

	err = data.UpdateTask(objectID, updateTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = data.DeleteTask(objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}