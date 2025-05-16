package todo

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
