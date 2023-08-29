// main.go
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

func main() {
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Todo List!")
	})

	e.GET("/todos", getTodos)
	e.POST("/todos", createTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	// Start the server
	e.Start(":8080")
}

func getTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func createTodo(c echo.Context) error {
	var newTodo Todo
	if err := c.Bind(&newTodo); err != nil {
		return err
	}
	newTodo.ID = len(todos) + 1
	newTodo.Status = "Incomplete"
	todos = append(todos, newTodo)
	return c.JSON(http.StatusCreated, newTodo)
}

func updateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var updatedTodo Todo
	if err := c.Bind(&updatedTodo); err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			return c.JSON(http.StatusOK, updatedTodo)
		}
	}

	return c.NoContent(http.StatusNotFound)
}

func deleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.JSON(http.StatusOK, echo.Map{
				"message": "Entry deleted successfully",
			})
		}
	}

	return c.NoContent(http.StatusNotFound)
}
