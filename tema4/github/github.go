package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://api.github.com/search/issues"

type Search struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Id     int
	Number int
	Title  string
	User   *User
}

type User struct {
	Login string
	Id    int
}

func GetIssues() (*Search, error) {
	response, err := http.Get(url + "?q=q")
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("status not 200")
	}

	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var jsonStr Search
	if err = json.Unmarshal(str, &jsonStr); err != nil {
		return nil, err
	}

	return &jsonStr, nil
}
