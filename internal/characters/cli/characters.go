package cli

import (
	"encoding/csv"
	"log"
	"os"

	characters "github.com/jvc9109/go-first-app/internal/characters"
	"github.com/jvc9109/go-first-app/internal/characters/fetching"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const (
	saveFileFlag = "filename"
	getAllFlag   = "all"
	pageFlag     = "page"
)

func InitCharactersCmd(service fetching.Service) *cobra.Command {
	characterCmd := &cobra.Command{
		Use:   "characters",
		Short: "Retrive characters data",
		Run:   runCharactersFn(service),
	}

	characterCmd.Flags().StringP(saveFileFlag, "f", "", "file where the result is stored")
	characterCmd.Flags().BoolP(getAllFlag, "a", false, "Required Recover all data through pages")
	characterCmd.Flags().StringP(pageFlag, "p", "", "Character Data whitin specific page")
	characterCmd.MarkFlagRequired(saveFileFlag)

	return characterCmd
}

func runCharactersFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		var chars []characters.Character
		var err error
		all, _ := cmd.Flags().GetBool(getAllFlag)
		page, _ := cmd.Flags().GetString(pageFlag)
		filename, _ := cmd.Flags().GetString(saveFileFlag)

		if all {
			chars, err = service.FetchAllCharacters()
		} else if page != "" {
			chars, err = service.FetchFromPageCharacters(page)
		} else {
			chars, err = service.FetchCharacters()
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
