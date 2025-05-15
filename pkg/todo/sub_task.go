package todo

// Entity representing a sub task.
type SubTask struct {
	Id     int
	Title  string
	Status TaskStatus
	Note   string
}
