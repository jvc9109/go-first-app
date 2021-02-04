package fetching

import (
	"github.com/jvc9109/go-first-app/internal/characters"
	"github.com/pkg/errors"
)

type Service interface {
	FetchCharacters() ([]characters.Character, error)
	FetchAllCharacters() ([]characters.Character, error)
	FetchFromPageCharacters(page string) ([]characters.Character, error)
	FetchByID(id int) (characters.Character, error)
}

type service struct {
	cR characters.CharacterRepo
}

func NewService(r characters.CharacterRepo) Service {
	return &service{r}
}

func (s *service) FetchCharacters() ([]characters.Character, error) {
	return s.cR.GetCharacters()
}

func (s *service) FetchAllCharacters() ([]characters.Character, error) {
	return s.cR.GetAllCharacters()
}

func (s *service) FetchFromPageCharacters(page string) ([]characters.Character, error) {
	return s.cR.GetCharactersFromPage(page)
}

func (s *service) FetchByID(id int) (characters.Character, error) {

	chars, err := s.cR.GetAllCharacters()

	if err != nil {
		return characters.Character{}, err
	}

	for _, chr := range chars {
		if chr.CharacterID == id {
			return chr, nil
		}
	}

	return characters.Character{}, errors.Errorf("Character with id %d not found", id)

}
