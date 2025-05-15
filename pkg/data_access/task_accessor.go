package data_access

import (
	"os"
	"strings"

	filehandling "github.com/Amonteverde04/TodoGo/internal/file_handling"
	"github.com/Amonteverde04/TodoGo/internal/reflection"
	"github.com/Amonteverde04/TodoGo/pkg/entity"
	"github.com/Amonteverde04/TodoGo/pkg/todo"
)

// Represents an accessor that accesses and updates tasks.
type TaskAccessor struct {
	file os.File
}

// Returns an instance of a TaskAccessor.
func NewTaskAccessor() TaskAccessor {
	file := filehandling.TryOpenFile("test.txt")

	return TaskAccessor{
		file: *file,
	}
}

// Gets all tasks.
func (taskAccessor TaskAccessor) GetAll() ([]*entity.TaskEntity, error) {
	return []*entity.TaskEntity{}, nil
}

// Gets a task by id.
func (taskAccessor TaskAccessor) GetById(id string) (*entity.TaskEntity, error) {
	return &entity.TaskEntity{}, nil
}

// Adds a task.
func (taskAccessor TaskAccessor) Add(taskData *todo.Task) (string, error) {
	if filehandling.FileIsEmpty(&taskAccessor.file) {
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
