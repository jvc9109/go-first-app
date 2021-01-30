package main

import (
	"log"

	"github.com/jvc9109/go-first-app/beers-cli/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	err := doc.GenMarkdownTree(rootCmd, "./doc")
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.Execute()
}
