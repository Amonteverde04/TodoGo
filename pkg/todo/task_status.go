package todo

// TaskStatus represents an int.
type TaskStatus int

// Representations of different status values.
const (
	InProgress TaskStatus = iota + 1
	ToReview
	OnHold
	Done
)

// TaskStatus dictionary.
var taskStatusName = map[TaskStatus]string{
	Done:       "DONE",
	InProgress: "IN PROGRESS",
	ToReview:   "TO REVIEW",
	OnHold:     "ON HOLD",
}

// Returns the string representation of a TaskStatus.
func TaskStatusToString(value TaskStatus) string {
	return taskStatusName[value]
}

// Returns true, if the integer value is in the taskStatusName dictionary.
func TaskStatusExists(value int) bool {
	_, exists := taskStatusName[TaskStatus(value)]
	return exists
}
