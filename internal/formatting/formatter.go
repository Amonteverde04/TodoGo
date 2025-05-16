package formatting

import (
	"encoding/json"

	"github.com/Amonteverde04/TodoGo/internal/error_handling"
)

// Converts some type to readable json.
func ToJSON(item any) string {
	b, err := json.MarshalIndent(item, "", "   ")
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	return string(b)
}
