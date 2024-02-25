package json

import (
	"encoding/json"
	"io"
)

// Parse parses the JSON-encoded body and stores the result in the value pointed to by v.
func Parse(body io.Reader, v any) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
