package data_access

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Amonteverde04/TodoGo/internal/error_handling"
	"github.com/Amonteverde04/TodoGo/internal/file_handling"
	"github.com/Amonteverde04/TodoGo/internal/reflection"
	"github.com/Amonteverde04/TodoGo/pkg/entity"
	"github.com/Amonteverde04/TodoGo/pkg/todo"
)

const (
	file_name = "task.csv"
)

// Represents an accessor that accesses and updates tasks.
type TaskAccessor struct {
	file os.File
}

// Returns an instance of a TaskAccessor.
func NewTaskAccessor() TaskAccessor {
	file := file_handling.TryOpenFile(file_name)

	return TaskAccessor{
		file: *file,
	}
}

// Gets all tasks.
func (taskAccessor TaskAccessor) GetAll() ([]entity.TaskEntity, error) {
	if file_handling.FileIsEmpty(&taskAccessor.file) {
		return []entity.TaskEntity{}, nil
	}

	reader := csv.NewReader(&taskAccessor.file)
	var tasks []entity.TaskEntity

	// Skip first header row.
	reader.Read()

	// Read the data rows.
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			error_handling.HandleError(err.Error(), 1)
		}

		// Create a new record and populate the fields.
		goalStatus, _ := strconv.Atoi(row[5])
		task := entity.TaskEntity{
			Entity: entity.Entity{
				Id:        row[0],
				CreatedAt: row[1],
				UpdatedAt: row[2],
			},
			Task: todo.Task{
				Title:      row[3],
				Goal:       row[4],
				GoalStatus: todo.TaskStatus(goalStatus),
				GoalNote:   row[6],
			},
		}

		// Append the record to the slice.
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Adds a task.
func (taskAccessor TaskAccessor) Add(taskData *todo.Task) (string, error) {
	if file_handling.FileIsEmpty(&taskAccessor.file) {
		InstantiateTaskFile(&taskAccessor.file)
	}

	id := WriteTaskData(&taskAccessor.file, *taskData)
	return id, nil
}

// Updates a task.
func (taskAccessor TaskAccessor) Update(id int, taskData *todo.Task) error {
	return nil
}

// Deletes a task.
func (taskAccessor TaskAccessor) Delete(id int) error {
	return nil
}

// Creates data store file for tasks in csv format.
func InstantiateTaskFile(file *os.File) {
	entityPropertyNameSlice := reflection.ReflectProperties(entity.Entity{})
	taskPropertyNameSlice := reflection.ReflectProperties(todo.Task{})
	combinedPropertyNameSlice := append(entityPropertyNameSlice, taskPropertyNameSlice...)
	file.WriteString(strings.Join(combinedPropertyNameSlice, ","))
}

// Appends task data in csv format.
func WriteTaskData(file *os.File, taskData todo.Task) string {
	taskToStore := CreateTaskEntity(taskData)
	entityPropertyValueSlice := reflection.ReflectValues(taskToStore.Entity)
	taskPropertyValueSlice := reflection.ReflectValues(taskToStore.Task)
	combinedPropertyValueSlice := append(entityPropertyValueSlice, taskPropertyValueSlice...)
	file.WriteString("\n" + strings.Join(combinedPropertyValueSlice, ","))
	return taskToStore.Id
}

// Creates a task entity to be stored.
func CreateTaskEntity(taskData todo.Task) entity.TaskEntity {
	return entity.NewTaskEntity(taskData)
}
