package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tarcea/go-api-todo/DB/data"
)

func GetTodos(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := list.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

func AddTodo(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo data.Item
		var body struct {
			Title       string
			Description string
			Done        bool
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error reading req.body"})
			return
		}
		currentTime := time.Now().String()
		todo = data.Item{
			ID:          uuid.New().String(),
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
			CreatedAt:   currentTime}

		list.Add(todo)

		c.JSON(200, todo)
	}
}

func GetTodoById(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo data.Item
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no id provided"})
			return
		}

		for _, item := range list.Items {
			if item.ID == id {
				todo = item
			}

		}

		if todo.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no item found with this ID"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(list *data.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo data.Item
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no id provided"})
			return
		}
		var pos int
		for _, item := range list.Items {
			if item.ID == id {
				todo = item
				list.Items = append(list.Items[:pos], list.Items[pos+1:]...)
				if pos > 0 {
					pos = pos - 1
				}
				continue
			}
			pos++
		}

		if todo.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no item found with this ID"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}
