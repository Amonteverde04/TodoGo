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
	"github.com/google/uuid"
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

// Gets a task using its id.
func (taskAccessor TaskAccessor) GetById(id uuid.UUID) (entity.TaskEntity, error) {
	if file_handling.FileIsEmpty(&taskAccessor.file) {
		return entity.TaskEntity{}, nil
	}

	reader := csv.NewReader(&taskAccessor.file)
	var taskToUpdate entity.TaskEntity

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
		if row[0] == id.String() {
			// Bind record data.
			goalStatus, _ := strconv.Atoi(row[5])
			taskToUpdate = entity.TaskEntity{
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
			break
		}
	}

	return taskToUpdate, nil
}

// Adds a task.
func (taskAccessor TaskAccessor) Add(taskData *todo.Task) (string, error) {
	if file_handling.FileIsEmpty(&taskAccessor.file) {
		InstantiateTaskFile(&taskAccessor.file)
	}

	id := WriteTaskData(&taskAccessor.file, *taskData)
	return id, nil
}

// Updates a task
func (taskAccessor TaskAccessor) Update(taskData entity.TaskEntity) error {
	if file_handling.FileIsEmpty(&taskAccessor.file) {
		return nil
	}

	reader := csv.NewReader(&taskAccessor.file)

	// Skip first header row.
	records, err := reader.ReadAll()
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	// Read the data rows and update.
	for i, row := range records {
		if i == 0 {
			continue
		}

		if row[0] == taskData.Entity.Id {
			// Update record data.
			row[2] = taskData.Entity.UpdatedAt
			row[3] = taskData.Task.Title
			row[4] = taskData.Task.Goal
			row[5] = strconv.Itoa(int(taskData.Task.GoalStatus))
			row[6] = taskData.Task.GoalNote
			break
		}
	}

	RewriteFile(records)
	return nil
}

// Deletes a task.
func (taskAccessor TaskAccessor) Delete(id uuid.UUID) error {
	if file_handling.FileIsEmpty(&taskAccessor.file) {
		return nil
	}

	reader := csv.NewReader(&taskAccessor.file)

	// Skip first header row.
	records, err := reader.ReadAll()
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	// Read the data rows and keep everything but removed item.
	var recordsToKeep [][]string
	for i, row := range records {
		if i == 0 || row[0] != id.String() {
			recordsToKeep = append(recordsToKeep, row)
		}
	}

	RewriteFile(recordsToKeep)
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

// Overwrites file records. Used after delete or update.
func RewriteFile(records [][]string) {
	// Reopen file to rewrite.
	file, err := os.Create(file_name)
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write updated records.
	err = writer.WriteAll(records)
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	// Remove trailing white space.
	trimmedContent := TrimWhiteSpace()
	err = os.WriteFile(file_name, []byte(trimmedContent), 0644)
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}
}

// Creates a task entity to be stored.
func CreateTaskEntity(taskData todo.Task) entity.TaskEntity {
	return entity.NewTaskEntity(taskData)
}

// Trim whitespace from update and delete.
func TrimWhiteSpace() string {
	content, err := os.ReadFile(file_name)
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}
	return strings.TrimRight(string(content), "\n")
}
