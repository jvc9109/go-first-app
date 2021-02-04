package main

import (
	"github.com/jvc9109/go-first-app/internal/characters/cli"
	ramapi "github.com/jvc9109/go-first-app/internal/characters/storage"
	"github.com/spf13/cobra"
)

func main() {
	apiRepo := ramapi.NewApiRepository()

	rootCmd := &cobra.Command{Use: "ram-cli"}
	rootCmd.AddCommand(cli.InitCharactersCmd(apiRepo))
	rootCmd.Execute()
}
