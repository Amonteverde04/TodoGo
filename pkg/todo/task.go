package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// Represents a task.
type Task struct {
	Title      string
	Goal       string
	GoalStatus TaskStatus
	GoalNote   string
}

// Returns a new task object.
func NewTask(title string, goal string, goalStatus int, goalNote string) Task {
	return Task{
		Title:      title,
		Goal:       goal,
		GoalStatus: TaskStatus(goalStatus),
		GoalNote:   goalNote,
	}
}

// Converts instance of task to readable json.
func (task Task) TaskToJson() string {
	b, err := json.MarshalIndent(task, "", "   ")
	if err != nil {
		fmt.Println("Error converting task to JSON.")
		os.Exit(int(2))
	}

	return string(b)
}
