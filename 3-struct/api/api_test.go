package api

import (
	"app/bin/config"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

const masterKey string = "$2a$10$SI7Jeohq4F1eAeX9DBbprOZmk5sJ1JskF8tqPBMFEO/0dWhwhsPuW"
const apiUrl = "https://api.jsonbin.io/v3"

type TestJson struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TestCreateBin(t *testing.T) {
	//Arrange
	conf := createTestConfig()
	api := NewApi(conf)
	binName := "testName"
	//Act
	bin, err := api.CreateBin(createTestJsonData("one", "two"), binName)
	//Assert
	if err != nil {
		t.Errorf("Can't create json bin because of %s", err.Error())
	}
	if bin.Metadata.Name != binName {
		t.Errorf("Expected result is not reached, Expected=%s, actual=%s", binName, bin.Metadata.Name)
	}
	defer func() {
		err := api.DeleteBin(bin.Metadata.Id)
		if err != nil {
			t.Fatalf("Can't clear test data for create method because of %s", err.Error())
		}
	}()
}

func TestUpdateBin(t *testing.T) {
	//Arrange
	conf := createTestConfig()
	api := NewApi(conf)
	binName := "testName"
	bin, err := api.CreateBin(createTestJsonData("one", "two"), binName)
	if err != nil {
		t.Errorf("Can't create json bin because of %s", err.Error())
	}
	defer func() {
		err := api.DeleteBin(bin.Metadata.Id)
		if err != nil {
			t.Fatalf("Can't clear test data for create method because of %s", err.Error())
		}
	}()
	expected := "three"
	//Act
	updated, err := api.UpdateBin(createTestJsonData(expected, "four"), bin.Metadata.Id)
	if err != nil {
		t.Errorf("Can't update json bin because of %s", err.Error())
	}
	//Assert
	if updated.Record["name"] != expected {
		t.Errorf("Expected result is not reached, Expected=%s, actual=%s", expected, updated.Record["name"])
	}
}

func TestGetBin(t *testing.T) {
	//Arrange
	conf := createTestConfig()
	api := NewApi(conf)
	binName := "testName"
	bin, err := api.CreateBin(createTestJsonData("one", "two"), binName)
	if err != nil {
		t.Errorf("Can't create json bin because of %s", err.Error())
	}
	defer func() {
		err := api.DeleteBin(bin.Metadata.Id)
		if err != nil {
			t.Fatalf("Can't clear test data for create method because of %s", err.Error())
		}
	}()
	expected := bin.Record["name"]
	//Act
	actual, err := api.ReadBin(bin.Metadata.Id)
	if err != nil {
		t.Errorf("Can't update json bin because of %s", err.Error())
	}
	//Assert
	if expected != actual.Record["name"] {
		t.Errorf("Expected result is not reached, Expected=%s, actual=%s", expected, actual.Record["name"])
	}
}

func TestDeleteBin(t *testing.T) {
	//Arrange
	conf := createTestConfig()
	api := NewApi(conf)
	binName := "testName"
	bin, err := api.CreateBin(createTestJsonData("one", "two"), binName)
	if err != nil {
		t.Errorf("Can't create json bin because of %s", err.Error())
	}
	//Act
	err = api.DeleteBin(bin.Metadata.Id)
	if err != nil {
		t.Errorf("Can't delete bin because of %s", err.Error())
	}
	_, err = api.ReadBin(bin.Metadata.Id)
	//Assert
	if err == nil || !strings.Contains(err.Error(), "404") {
		t.Error("Expected result not reached, bin must be deleted but it is not")
	}
}

func createTestConfig() *config.Config {
	if apiUrl == "" {
		panic("Api url(API_URL) is not found in env file")
	}
	u, err := url.Parse(apiUrl)
	if err != nil {
		panic(fmt.Sprintf("Can not parse api url(API_URL) to url string %v", err))
	}
	config := config.Config{
		Key:    masterKey,
		ApiUrl: *u,
	}
	return &config
}

func createTestJsonData(name string, value string) []byte {
	data := TestJson{name, value}
	json, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("Can't map data to json bytes %s", err.Error()))
	}
	return json
}
