package validator

import (
	"github.com/Amonteverde04/TodoGo/internal/error_handling"
	"github.com/Amonteverde04/TodoGo/pkg/todo"
)

// Validates a task name and status.
func ValidateTaskTitleAndStatus(title string, status int) {
	if len(title) < 1 {
		error_handling.HandleError("The title argument is required when using the add subcommand.", 1)
	}

	if !StatusInputIsValid(status) {
		error_handling.HandleError("The goal status argument can only be 1 (IN PROGRESS), 2 (TO REVIEW), 3 (ON HOLD), or 4 (DONE).", 1)
	}
}

// Validates that a user input a valid status for goals and tasks.
func StatusInputIsValid(value int) bool {
	if value == 0 {
		return true
	}

	return todo.TaskStatusExists(value)
}
