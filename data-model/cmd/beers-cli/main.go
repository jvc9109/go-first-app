package main

import (
	"github.com/jvc9109/go-first-app/data-model/internal/cli"
	"github.com/jvc9109/go-first-app/data-model/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {
	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	rootCmd.Execute()
}
