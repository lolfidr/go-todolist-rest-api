package main

import (
	"fmt"
	"todolist/http"
	"todolist/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlres(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed start https server", err)
	}
}
