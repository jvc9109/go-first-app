package cli

import (
	"fmt"
	"github.com/spf13/cobra"
) 

type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"a": "cervesa1",
	"b": "cerversa2",
}

const idFlag = "id"

func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{
		Use: "beers",
		Short: "Print data about beers",
		Run: runBeersFn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn() CobraFn{

	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}

}