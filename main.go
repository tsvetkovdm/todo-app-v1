package main

import (
	"todo-app/scanner"
	"todo-app/todo"
)

func main() {
	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Start()

}
