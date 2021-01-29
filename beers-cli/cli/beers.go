package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "Mad Jack Mixer",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "Keystone Ice",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "Belgian Moon",
}

const idBeerFlag = "id"

func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	beersCmd.Flags().StringP(idBeerFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn() CobraFn {

	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idBeerFlag)

		if id != "" {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}

}
