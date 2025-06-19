package todo

import "strings"

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
		Title:      SanitizeCommas(title),
		Goal:       SanitizeCommas(goal),
		GoalStatus: TaskStatus(goalStatus),
		GoalNote:   SanitizeCommas(goalNote),
	}
}

func SanitizeCommas(value string) string {
	return strings.ReplaceAll(value, ",", " ")
}
