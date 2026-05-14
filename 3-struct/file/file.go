package file

import (
	"errors"
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

func IsJson(name string) bool {
	return strings.HasSuffix(name, ".json")
}

func WriteFile(name string, data []byte) error {
	if !IsJson(name) {
		return errors.New("File can not be written because it is not json file")
	}
	return os.WriteFile(name, data, 0644)
}
