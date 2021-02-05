package fetching_test 

import (
	"errors"
	"testing"
	. "github.com/jvc9109/go-first-app/internal/characters/fetching"
	"github.com/jvc9109/go-first-app/internal/characters/storage/inmem"
)

func TestFetchByID(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
		err   error
	}{
		"valid character":     {input: 5, want: 5, err: nil},
		"not found character": {input: 1000000, err: errors.New("error")},
	}

	repo := inmem.NewRepository()
	service := NewService(repo)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			chr, err := service.FetchByID(tc.input)
			if err != nil && tc.err == nil {
				t.Fatalf("not expected erros. Got %v", err)
			}

			if err == nil && tc.err != nil {
				t.Error("expected error and got nil")
			}

			if chr.CharacterID != tc.want {
				t.Fatalf("expected %d, got: %d", tc.want, chr.CharacterID)
			}

		})
	}
}
