package validator

import (
	"github.com/Amonteverde04/TodoGo/internal/error_handling"
	"github.com/Amonteverde04/TodoGo/pkg/todo"
	"github.com/google/uuid"
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

// Validate a task update. If ALL fields are empty or a default value, throw error.
func ValidateTaskUpdate(updateTitle string, updateGoal string, updateGoalStatus int, updateGoalNote string) {
	if len(updateTitle) < 1 && len(updateGoal) < 1 && len(updateGoalNote) < 1 && !StatusInputIsValid(updateGoalStatus) {
		error_handling.HandleError("At least one argument aside from id is required when using the update subcommand.", 1)
	}
}

// Validates a task id.
func ValidateId(id string) {
	if len(id) < 1 {
		error_handling.HandleError("The id argument is required when using the remove or update subcommands.", 1)
	}

	if !IdIsGuid(id) {
		error_handling.HandleError("The id argument must be a valid GUID when using the remove or update subcommands.", 1)
	}
}

// Validates that a user input a valid status for goals and tasks.
func StatusInputIsValid(value int) bool {
	if value == 0 {
		return false
	}

	return todo.TaskStatusExists(value)
}

func IdIsGuid(id string) bool {
	err := uuid.Validate(id)
	return !(err != nil)
}
