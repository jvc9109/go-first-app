package cli

import (
	"encoding/csv"
	"log"
	"os"

	characters "github.com/jvc9109/go-first-app/rickmorty-api/internal/characters"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const (
	saveFileFlag = "filename"
	getAllFlag   = "all"
	pageFlag     = "page"
)

func InitCharactersCmd(repository characters.CharacterRepo) *cobra.Command {
	characterCmd := &cobra.Command{
		Use:   "characters",
		Short: "Retrive characters data",
		Run:   runCharactersFn(repository),
	}

	characterCmd.Flags().StringP(saveFileFlag, "f", "", "file where the result is stored")
	characterCmd.Flags().BoolP(getAllFlag, "a", false, "Required Recover all data through pages")
	characterCmd.Flags().StringP(pageFlag, "p", "", "Character Data whitin specific page")
	characterCmd.MarkFlagRequired(saveFileFlag)

	return characterCmd
}

func runCharactersFn(repository characters.CharacterRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		var chars []characters.Character
		var err error
		all, _ := cmd.Flags().GetBool(getAllFlag)
		page, _ := cmd.Flags().GetString(pageFlag)
		filename, _ := cmd.Flags().GetString(saveFileFlag)

		if all {
			chars, err = repository.GetAllCharacters()
		} else if page != "" {
			chars, err = repository.GetCharactersFromPage(page)
		} else {
			chars, err = repository.GetCharacters()
		}

		if err != nil {
			log.Fatal(err.Error())
		}


		file, err := os.Create(filename)

		if err != nil {
			log.Fatal(err.Error())
		}

		writer := csv.NewWriter(file)
		headers := chars[0].GetHeaders()

		writer.Write(headers)

		for _, chr := range chars {
			values := chr.ToSlice()
			writer.Write(values)

		}

		writer.Flush()
		file.Close()

	}
}
