package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Amonteverde04/TodoGo/internal/error_handling"
	"github.com/Amonteverde04/TodoGo/internal/validator"
	"github.com/Amonteverde04/TodoGo/pkg/data_access"
	"github.com/Amonteverde04/TodoGo/pkg/entity"
	"github.com/Amonteverde04/TodoGo/pkg/todo"
)

func main() {
	addCmd := flag.NewFlagSet("addTask", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Sets the title of the task being added.")
	addGoal := addCmd.String("goal", "", "Sets the high level goal of the task being added.")
	addGoalStatus := addCmd.Int("goalStatus", 0, "Sets the status of the high level goal. Available options: 1 (IN PROGRESS), 2 (TO REVIEW), 3 (ON HOLD), or 4 (DONE).")
	addGoalNote := addCmd.String("goalNote", "", "Sets notes about the high level goal.")

	//updateCmd := flag.NewFlagSet("updateTask", flag.ExitOnError)
	//updateId := updateCmd.Int("id", 0, "Reference to the task that should be updated.")
	//updateTitle := updateCmd.String("title", "", "Sets the title of the task being added.")
	//updateGoal := updateCmd.String("goal", "", "Sets the high level goal of the task being added.")
	//updateGoalStatus := updateCmd.Int("goalStatus", 0, "Sets the status of the high level goal. Available options: 1 (IN PROGRESS), 2 (TO REVIEW), 3 (ON HOLD), or 4 (DONE).")
	//updateGoalNote := updateCmd.String("goalNote", "", "Sets notes about the high level goal.")

	//addSubTask := addCmd.String("subTask", "", "Sets the sub task.")
	//addSubTaskStatus := addCmd.Int("subTaskStatus", 0, "Sets the status of the sub task.")
	//addSubTaskNote := addCmd.String("subTaskNote", "", "Sets notes about the sub task.")

	if len(os.Args) < 2 {
		error_handling.HandleError("expected 'add' subcommands", 1)
	}

	switch os.Args[1] {
	case "addTask":
		// Parse arguments.
		addCmd.Parse(os.Args[2:])
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
		fmt.Print(taskToAdd.TaskToJson())
	case "updateTask":
		//updateCmd.Parse(os.Args[2:])
		//if *updateId == 0 {
		//	handleError("The id argument is required when using the update subcommand.", 1)
		//}
		//validateTaskTitleAndStatus(*updateTitle, *updateGoalStatus)
	default:
		error_handling.HandleError("expected 'add' subcommands", 1)
	}
}
