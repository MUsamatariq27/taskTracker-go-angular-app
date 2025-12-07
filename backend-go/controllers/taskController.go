package controllers

import (
	"net/http"

	"github.com/MUsamaT/task-tracker/database"
	"github.com/MUsamaT/task-tracker/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(ctx *gin.Context) {

	var input models.Task
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input!",
		})
		return
	}

	query := `INSERT INTO tasks (title, description, completed, user_id) VALUES ($1, $2, $3, $4)`
	_, err = database.DB.Exec(query, input.Title, input.Description, input.Completed, input.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating task!"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created succesfully!"})

}

func GetAllTasks(ctx *gin.Context) {

	query := `SELECT id, title, description, completed, user_id FROM tasks`
	rows, err := database.DB.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Fetching Tasks!"})
		return
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error scanning task record!",
			})
			return
		}
		tasks = append(tasks, task)

	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func GetUserTasks(ctx *gin.Context) {

	id := ctx.Param("id")

	var input models.Task
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input!",
		})
		return
	}

	rows, err := database.DB.Query("SELECT id, title, description, completed, user_id FROM tasks WHERE user_id=$1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tasks"})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.UserID); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning task"})
			return
		}
		tasks = append(tasks, task)
	}
	if tasks == nil {
		tasks = []models.Task{}
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func UpddateTask(ctx *gin.Context) {

	id := ctx.Param("id")

	var input models.Task
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input!",
		})
		return
	}

	query := `UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4`
	result, err := database.DB.Exec(query, input.Title, input.Description, input.Completed, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error updating task!",
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking update result!",
		})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully!",
	})
}

func DeleTask(ctx *gin.Context) {

	id := ctx.Param("id")

	query := `DELETE FROM tasks WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error deleting task!",
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking delete result!",
		})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully!",
	})
}
