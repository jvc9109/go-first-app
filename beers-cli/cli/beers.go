package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

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
