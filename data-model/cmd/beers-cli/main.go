package main

import (
	"github.com/jvc9109/go-first-app/data-model/internal/cli"
	"github.com/jvc9109/go-first-app/data-model/internal/storage/ontario"
	"github.com/spf13/cobra"
)

func main() {
	ontarioRepo := ontario.NewOntarioBeerRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(ontarioRepo))
	rootCmd.Execute()
}
