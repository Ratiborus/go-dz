package api

import (
	"app/bin/bins"
	"app/bin/config"
	"app/bin/mapper"
	"bytes"
	"fmt"
	"net/http"
)

const (
	key         = "X-Master-Key"
	name        = "X-Bin-Name"
	contentType = "Content-Type"
)

type Api struct {
	config *config.Config
}

func NewApi(conf *config.Config) *Api {
	return &Api{config: conf}
}

func (api *Api) CreateBin(body []byte, binName string) (*bins.Bin, error) {
	link := api.config.BinUrl()
	r, err := http.NewRequest(http.MethodPost, link.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.Header.Add(key, api.config.Key)
	r.Header.Add(name, binName)
	r.Header.Add(contentType, "application/json")
	res, err := makeRequest(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Create bin - status code is %v", res.StatusCode)
	}
	bin, err := mapper.Map(res.Body)
	if err != nil {
		return nil, err
	}
	return bin, nil
}

func (api *Api) UpdateBin(body []byte, binId string) (*bins.Bin, error) {
	link := api.config.BinUrl().JoinPath(binId)
	r, err := http.NewRequest(http.MethodPut, link.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.Header.Add(key, api.config.Key)
	r.Header.Add(contentType, "application/json")
	res, err := makeRequest(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Update bin - status code is %v", res.StatusCode)
	}
	bin, err := mapper.Map(res.Body)
	if err != nil {
		return nil, err
	}
	return bin, nil
}

func (api *Api) ReadBin(binId string) (*bins.Bin, error) {
	link := api.config.BinUrl().JoinPath(binId)
	r, err := http.NewRequest(http.MethodGet, link.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add(key, api.config.Key)
	r.Header.Add(contentType, "application/json")
	res, err := makeRequest(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Read bin - status code is %v", res.StatusCode)
	}
	bin, err := mapper.Map(res.Body)
	if err != nil {
		return nil, err
	}
	return bin, nil
}

func (api *Api) DeleteBin(binId string) error {
	link := api.config.BinUrl().JoinPath(binId)
	r, err := http.NewRequest(http.MethodDelete, link.String(), nil)
	if err != nil {
		return err
	}
	r.Header.Add(key, api.config.Key)
	res, err := makeRequest(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 && res.StatusCode != 404 {
		return fmt.Errorf("Delete bin - status code is %v", res.StatusCode)
	}
	return nil
}

func makeRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
