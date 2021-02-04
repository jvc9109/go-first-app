package fetching

import (
	"testing"

	"github.com/jvc9109/go-first-app/internal/characters/storage/inmem"
)

func TestFetchByID(t *testing.T) {
	repo := inmem.NewRepository()

	service := NewService(repo)

	expected := 5

	c, err := service.FetchByID(expected)

	if err != nil {
		t.Fatalf("expected %d, fot an error %v", expected, err)
	}

	if c.CharacterID != expected {
		t.Fatalf("expected %d, got: %d ", expected, c.CharacterID)
	}

}
