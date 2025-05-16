package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Amonteverde04/TodoGo/internal/error_handling"
	"github.com/Amonteverde04/TodoGo/internal/formatting"
	"github.com/Amonteverde04/TodoGo/internal/validator"
	"github.com/Amonteverde04/TodoGo/pkg/data_access"
	"github.com/Amonteverde04/TodoGo/pkg/entity"
	"github.com/Amonteverde04/TodoGo/pkg/todo"
)

func main() {
	// List command.
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	// AddTask command and corresponding arguments.
	addTaskCmd := flag.NewFlagSet("addTask", flag.ExitOnError)
	addTitle := addTaskCmd.String("title", "", "Sets the title of the task being added.")
	addGoal := addTaskCmd.String("goal", "", "Sets the high level goal of the task being added.")
	addGoalStatus := addTaskCmd.Int("goalStatus", 0, "Sets the status of the high level goal of the task being added. Available options: 1 (IN PROGRESS), 2 (TO REVIEW), 3 (ON HOLD), or 4 (DONE).")
	addGoalNote := addTaskCmd.String("goalNote", "", "Sets notes about the high level goal of the task being added.")
	// UpdateTask command and corresponding arguments.
	updateTaskCmd := flag.NewFlagSet("updateTask", flag.ExitOnError)
	updateId := updateTaskCmd.String("id", "", "Sets the id, referencing the task that should be updated.")
	updateTitle := updateTaskCmd.String("title", "", "Sets the title of the task being updated.")
	updateGoal := updateTaskCmd.String("goal", "", "Sets the high level goal of the task being updated.")
	updateGoalStatus := updateTaskCmd.Int("goalStatus", 0, "Sets the status of the high level goal of the task being updated. Available options: 1 (IN PROGRESS), 2 (TO REVIEW), 3 (ON HOLD), or 4 (DONE).")
	updateGoalNote := updateTaskCmd.String("goalNote", "", "Sets notes about the high level goal of the task being updated.")
	// RemoveTask command and corresponding arguments.
	removeTaskCmd := flag.NewFlagSet("removeTask", flag.ExitOnError)
	removeTaskId := removeTaskCmd.String("id", "", "Sets the id, referencing the task that should be removed.")

	//addSubTask := addCmd.String("subTask", "", "Sets the sub task.")
	//addSubTaskStatus := addCmd.Int("subTaskStatus", 0, "Sets the status of the sub task.")
	//addSubTaskNote := addCmd.String("subTaskNote", "", "Sets notes about the sub task.")

	if len(os.Args) < 2 {
		error_handling.HandleInvalidSelection()
	}

	switch os.Args[1] {
	case "addTask":
		// Parse arguments.
		addTaskCmd.Parse(os.Args[2:])
		// Validate input.
		validator.ValidateTaskTitleAndStatus(*addTitle, *addGoalStatus)
		// Create task object that we will save.
		taskToAdd := todo.NewTask(*addTitle, *addGoal, *addGoalStatus, *addGoalNote)
		// Call implementation of data accessor's add method.
		_, err := data_access.DataAccessor[entity.TaskEntity, todo.Task].Add(data_access.NewTaskAccessor(), &taskToAdd)
		// Escape if error storing data.
		if err != nil {
			error_handling.HandleError("Could not add task. An error occured.", 1)
		}
		// Output stored data.
		fmt.Println("Task added:")
		fmt.Print(formatting.ToJSON(taskToAdd))
	case "updateTask":
		// Parse arguments.
		updateTaskCmd.Parse(os.Args[2:])
		// Validate input.
		validator.ValidateId(*updateId)
		validator.ValidateTaskUpdate(*updateTitle, *updateGoal, *updateGoalStatus, *updateGoalNote)
		// Parse validated GUID.
		id := formatting.ToGUID(*updateId)
		// Attempt get task from data source.
		taskToUpdate, getErr := data_access.DataAccessor[entity.TaskEntity, todo.Task].GetById(data_access.NewTaskAccessor(), id)
		// Escape if error getting data.
		if getErr != nil {
			error_handling.HandleError(getErr.Error(), 1)
		}
		// Replace data.
		taskToUpdate = UpdateTaskValues(taskToUpdate, *updateTitle, *updateGoal, *updateGoalStatus, *updateGoalNote)
		// Update data.
		updateErr := data_access.DataAccessor[entity.TaskEntity, todo.Task].Update(data_access.NewTaskAccessor(), taskToUpdate)
		// Escape if error updating data.
		if updateErr != nil {
			error_handling.HandleError(updateErr.Error(), 1)
		}
		// Output success.
		fmt.Println("Successfully updated task: ")
		fmt.Println(formatting.ToJSON(taskToUpdate))
	case "list":
		// Skip parsing of arguments.
		listCmd.Parse(nil)
		// Attempt get all from data source.
		tasks, err := data_access.DataAccessor[entity.TaskEntity, todo.Task].GetAll(data_access.NewTaskAccessor())
		// Escape if error getting data.
		if err != nil {
			error_handling.HandleError(err.Error(), 1)
		}
		// Output data.
		fmt.Println(formatting.ToJSON(tasks))
	case "removeTask":
		// Parse arguments.
		removeTaskCmd.Parse(os.Args[2:])
		// Validate input.
		validator.ValidateId(*removeTaskId)
		// Parse validated GUID.
		id := formatting.ToGUID(*removeTaskId)
		// Attempt remove from data source.
		removeErr := data_access.DataAccessor[entity.TaskEntity, todo.Task].Delete(data_access.NewTaskAccessor(), id)
		// Escape if error removing data.
		if removeErr != nil {
			error_handling.HandleError(removeErr.Error(), 1)
		}
		// Output success.
		fmt.Println("Successfully removed task with id of: " + id.String())
	default:
		error_handling.HandleInvalidSelection()
	}
}

// Updates the task property that changed and updates the updated time.
func UpdateTaskValues(taskToUpdate entity.TaskEntity, updateTitle string, updateGoal string, updateGoalStatus int, updateGoalNote string) entity.TaskEntity {
	if len(updateTitle) > 0 && taskToUpdate.Task.Title != updateTitle {
		taskToUpdate.Task.Title = updateTitle
	}

	if len(updateGoal) > 0 && taskToUpdate.Task.Goal != updateGoal {
		taskToUpdate.Task.Goal = updateGoal
	}

	if taskToUpdate.Task.GoalStatus != todo.TaskStatus(updateGoalStatus) {
		taskToUpdate.Task.GoalStatus = todo.TaskStatus(updateGoalStatus)
	}

	if len(updateGoalNote) > 0 && taskToUpdate.Task.GoalNote != updateGoalNote {
		taskToUpdate.Task.GoalNote = updateGoalNote
	}

	taskToUpdate.Entity.UpdatedAt = time.Now().UTC().String()
	return taskToUpdate
}
