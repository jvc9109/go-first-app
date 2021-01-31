package ramapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jvc9109/go-first-app/rickmorty-api/internal/characters"
)

const (
	charachtersEndpoint = "/character"
	apiUrl              = "https://rickandmortyapi.com/api"
)

type characterRepo struct {
	url string
}

func NewApiRepository() characters.CharacterRepo {
	return &characterRepo{url: apiUrl}
}

func (c *characterRepo) GetAllCharacters() (chars []characters.Character, err error) {
	var apiResponse characters.CharacterApi

	response, err := http.Get(fmt.Sprintf("%v%v", c.url, charachtersEndpoint))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(contents, &chars)
	if err != nil {
		return nil, err
	}

	chars = apiResponse.Results
	return
}

func (c *characterRepo) GetCharacters() (chars []characters.Character, err error) {
	var apiResponse characters.CharacterApi

	response, err := http.Get(fmt.Sprintf("%v%v", c.url, charachtersEndpoint))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}

	chars = apiResponse.Results

	return
}

func (c *characterRepo) GetCharactersFromPage(page string) (chars []characters.Character, err error) {
	var apiResponse characters.CharacterApi

	response, err := http.Get(fmt.Sprintf("%v%v?page=%v", c.url, charachtersEndpoint, page))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}

	chars = apiResponse.Results

	return
}
