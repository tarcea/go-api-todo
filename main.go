package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-api-todo/DB/data"
	"github.com/tarcea/go-api-todo/routes"
)

var list = data.New()

func init() {
	list.Add(data.Item{
		ID:          "ccb1af3c-6038-4f1b-cd27-3e0d99e964eb",
		Title:       "buy something",
		Description: "NOW",
		Done:        false,
		CreatedAt:   "2022-09-16 14:03:34.748452 +0200 CEST m=+114.141099209"})
	list.Add(data.Item{
		ID:          "ccb1af3c-6038-4f1b-1227-3e0d99e964eb",
		Title:       "walk the dog",
		Description: "grab the pet and walk it arount the yard",
		Done:        false,
		CreatedAt:   "2022-09-16 14:03:39.748452 +0200 CEST m=+114.141099209"})
}

func main() {
	r := gin.Default()

	r.GET("/", routes.Home())
	r.GET("/todos", routes.GetTodos(list))
	r.GET("/todos/:id", routes.GetTodoById(list))
	r.POST("/todos", routes.AddTodo(list))
	r.DELETE("/todos/:id", routes.DeleteTodo(list))

	// server listen on 8080 if no PORT specified in .env
	r.Run()

}
