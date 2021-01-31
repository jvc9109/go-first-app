package cli

import (
	characters "github.com/jvc9109/go-first-app/rickmorty-api/internal/characters"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const (
	saveFileFlag = "filename"
	getAllFlag = "all"
	pageFlag = "page"
)

func InitCharactersCmd(repository characters.CharacterRepo) *cobra.Command {
	characterCmd := &cobra.Command{
		Use: "characters",
		Short: "Retrive characters data",
	}
	
	characterCmd.Flags().StringP(saveFileFlag, "f", "", "file where the result is stored")
	characterCmd.Flags().StringP(getAllFlag, "a", "", "Recover all data through pages")
	characterCmd.Flags().StringP(pageFlag, "p", "", "Character Data whitin specific page")

	return characterCmd
}

func runCharactersFn(repository characters.CharacterRepo) CobraFn {
	return func (cmd *cobra.Command, args []string)  {
		chars, _ := repository.GetCharacters()
	}
}