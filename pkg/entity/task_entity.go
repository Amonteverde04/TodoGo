package entity

import (
	"time"

	"github.com/Amonteverde04/TodoGo/pkg/todo"
	"github.com/google/uuid"
)

// Represents a task entity for the database.
type TaskEntity struct {
	Entity
	todo.Task
}

// Returns a new task entity based on a task.
func NewTaskEntity(taskData todo.Task) TaskEntity {
	return TaskEntity{
		Entity: Entity{
			Id:        uuid.New().String(),
			CreatedAt: time.Now().UTC().String(),
			UpdatedAt: time.Now().UTC().String(),
		},
		Task: taskData,
	}
}
