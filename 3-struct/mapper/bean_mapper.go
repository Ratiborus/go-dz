package mapper

import (
	"app/bin/bins"
	"encoding/json"
	"io"
)

func Map(body io.ReadCloser) (*bins.Bin, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var bin bins.Bin
	err = json.Unmarshal(b, &bin)
	if err != nil {
		return nil, err
	}
	return &bin, nil
}
