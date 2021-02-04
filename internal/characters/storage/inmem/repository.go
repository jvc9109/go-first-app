package inmem

import (
	"github.com/jvc9109/go-first-app/internal/characters"
)

type repository struct {
}

func NewRepository() characters.CharacterRepo {
	return &repository{}
}

func (c *repository) GetAllCharacters() (chars []characters.Character, err error) {
	return c.setData(), nil
}

func (c *repository) GetCharacters() (chars []characters.Character, err error) {
	return
}

func (c *repository) GetCharactersFromPage(page string) (chars []characters.Character, err error) {
	return
}

func (r *repository) setData() []characters.Character {

	return []characters.Character{
		characters.NewCharacter(
			1, "Rick Sanchez", "Alive", "Human", "", "https://rickandmortyapi.com/api/character/avatar/1.jpeg", []string{
				"https://rickandmortyapi.com/api/episode/1",
				"https://rickandmortyapi.com/api/episode/2",
				"https://rickandmortyapi.com/api/episode/3",
			}, characters.NewGenderType("Male"),
		),
		characters.NewCharacter(
			5, "Jerry Smith", "Alive", "Human", "", "https://rickandmortyapi.com/api/character/avatar/5.jpeg", []string{
				"https://rickandmortyapi.com/api/episode/14",
				"https://rickandmortyapi.com/api/episode/12",
				"https://rickandmortyapi.com/api/episode/40",
			}, characters.NewGenderType("Male"),
		),
	}
}
