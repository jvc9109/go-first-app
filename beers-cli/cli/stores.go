package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

const idStoreFlag = "id"

func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Print data about stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(idBeerFlag, "i", "", "id of the store")

	return storesCmd
}

func runStoresFn() CobraFn {

	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idStoreFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}

}
