package entity

import "github.com/Amonteverde04/TodoGo/pkg/todo"

// Represents a task entity for the database.
type TaskEntity struct {
	Entity
	todo.Task
}
