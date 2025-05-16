package formatting

import (
	"encoding/json"

	"github.com/Amonteverde04/TodoGo/internal/error_handling"
	"github.com/google/uuid"
)

// Converts some type to readable json.
func ToJSON(item any) string {
	b, err := json.MarshalIndent(item, "", "   ")
	if err != nil {
		error_handling.HandleError(err.Error(), 1)
	}

	return string(b)
}

// Converts some type to a GUID.
func ToGUID(item string) uuid.UUID {
	guid, guidErr := uuid.Parse(item)
	// Escape if error parsing.
	if guidErr != nil {
		error_handling.HandleError(guidErr.Error(), 1)
	}

	return guid
}
