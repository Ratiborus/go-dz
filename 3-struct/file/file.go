package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("Can not read file %w", err)
	}
	return b, nil
}

func isJson(name string) bool {
	return strings.HasSuffix(name, ".json")
}
