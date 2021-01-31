package characters

import (
	"encoding/json"
	"fmt"
)

type Character struct {
	CharacterID   int         `json:"id"`
	Name          string      `json:"name"`
	Status        string      `json:"status"`
	Species       string      `json:"species"`
	CharacterType string      `json:"type"`
	Gender        *GenderType `json:"gender"`
	Image         string      `json:"image"`
	Episodes      []string    `json:"episode"`
}

func (c *Character) toSlice() (s []string) {

	s = []string{
		fmt.Sprint(c.CharacterID), c.Name, c.Status, c.Species, c.CharacterType, c.Gender.String(), c.Image,
	}

	return
}

type CharacterApi struct {
	Info    InfoApi     `json:"info"`
	Results []Character `json:"results"`
}

type InfoApi struct {
	Count    int    `json:"count"`
	Pages    int    `json:"pages"`
	NextPage string `json:"next"`
	PrevPage string `json:"prev"`
}

type GenderType int

const (
	Unknown GenderType = iota
	Male
	Female
	Genderless
)

func (g GenderType) String() string {
	return toString[g]
}

func NewGenderType(t string) *GenderType {
	genderType := toID[t]
	return &genderType
}

var toString = map[GenderType]string{
	Female:     "Female",
	Male:       "Male",
	Genderless: "Genderless",
	Unknown:    "unknown",
}

var toID = map[string]GenderType{
	"Female":     Female,
	"Male":       Male,
	"Genderless": Genderless,
	"unknown":    Unknown,
}

func (t *GenderType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t = toID[j]
	return nil
}

type CharacterRepo interface {
	GetCharacters() ([]Character, error)
	GetAllCharacters() ([]Character, error)
	GetCharactersFromPage(page string) ([]Character, error)
}

func NewCharacter(characterID int, name, status, species, characterType, image string, episodes []string, genderType *GenderType) (c Character) {
	c = Character{
		CharacterID:   characterID,
		Name:          name,
		Status:        status,
		Species:       species,
		CharacterType: characterType,
		Gender:        genderType,
		Image:         image,
		Episodes:      episodes,
	}
	return
}
