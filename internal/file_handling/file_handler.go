package file_handling

import (
	"os"

	"github.com/Amonteverde04/TodoGo/internal/error_handling"
)

// Attempt to open a file and if you can't, create it.
func TryOpenFile(fileName string) *os.File {
	// 0644: "you can read and write the file or directory and other users can only read it"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}
	return file
}

// Checks if a file is empty. Returns true if it is.
func FileIsEmpty(file *os.File) bool {
	fileInfo, err := file.Stat()

	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	return fileInfo.Size() == 0
}
