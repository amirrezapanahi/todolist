package main

import (
	"flag"
	"github.com/amirrezapanahi/todolist"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")

	flag.Parse()

	todos := &todo.Todos
}