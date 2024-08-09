package main

import (
	"net/http"
	"github.com/gin-conic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IntendedJSON(http.StatusOk, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
