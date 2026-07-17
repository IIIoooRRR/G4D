package main

import (
	"fmt"
	cmd2 "g4d-cli/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "g4d",
		Short: "G4D CLI",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Initialize a new G4D project",
		Run: func(cmd *cobra.Command, args []string) {
			cmd2.InitDir()
			fmt.Print("Initialized a new G4D project\n")
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
