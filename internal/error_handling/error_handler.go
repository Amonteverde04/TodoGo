package error_handling

import (
	"fmt"
	"os"
)

// Handles errors so we do not have to keep writing contents of this method over and over.
func HandleError(errorMessage string, errorCode int8) {
	fmt.Println(errorMessage)
	os.Exit(int(errorCode))
}
