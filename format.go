package gosta

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSON returns converted Go interface into a string format of JSON
func (*Client) JSON(m map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(m); err != nil {
		return "", fmt.Errorf("Error encoding query: %s", err)
	}

	return buf.String(), nil
}
