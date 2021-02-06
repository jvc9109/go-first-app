package ramapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jvc9109/go-first-app/internal/characters"
	"github.com/jvc9109/go-first-app/internal/characters/errors"
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
		return nil, errors.WrapDataUnreacheable(err, "error getting response from %s", charachtersEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading response from %s", charachtersEndpoint)
	}

	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "cant parse response into characters")
	}

	chars = apiResponse.Results
	numberPages := apiResponse.Info.Pages

	chr := make(chan []characters.Character)

	for i := 2; i <= numberPages; i++ {

		nextPage := i
		

		go func(chr chan []characters.Character, nextPage int) {
			var tempResult characters.CharacterApi
			response, _ := http.Get(fmt.Sprintf("%v%v?page=%v", c.url, charachtersEndpoint, nextPage))
			contents, _ := ioutil.ReadAll(response.Body)
			_ = json.Unmarshal(contents, &tempResult)
			chr <- tempResult.Results
		}(chr, nextPage)
		// if err != nil {
		// 	return nil, errors.WrapDataUnreacheable(err, "error reading response from %s", charachtersEndpoint)
		// }
		// if err != nil {
		// 	return nil, errors.WrapDataUnreacheable(err, "cant parse response into characters")
		// }
		// chars = append(chars, tempResult.Results...)
	}

	i := 2
	for i < numberPages {
		select {
		case chchrs := <-chr:
			chars = append(chars, chchrs...)
			i++
		}
	}

	return
}

func (c *characterRepo) GetCharacters() (chars []characters.Character, err error) {
	var apiResponse characters.CharacterApi

	response, err := http.Get(fmt.Sprintf("%v%v", c.url, charachtersEndpoint))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response from %s", charachtersEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading response from %s", charachtersEndpoint)
	}

	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "cant parse response into characters")
	}

	chars = apiResponse.Results

	return
}

func (c *characterRepo) GetCharactersFromPage(page string) (chars []characters.Character, err error) {
	var apiResponse characters.CharacterApi

	response, err := http.Get(fmt.Sprintf("%v%v?page=%v", c.url, charachtersEndpoint, page))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response from %s", charachtersEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading response from %s", charachtersEndpoint)
	}

	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "cant parse response into characters")
	}

	chars = apiResponse.Results

	return
}
