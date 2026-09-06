package main

import (
	"fmt"
	cmd2 "g4d-cli/cmd"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/analysis/singlechecker"
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
			fmt.Println("Initialized a new G4D project")
		}},
		&cobra.Command{
			Use:   "lint [path...]",
			Short: "Lint a G4D project",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					args = []string{"./..."}
				}
				oldArgs := os.Args
				os.Args = append([]string{"g4d"}, args...)
				defer func() { os.Args = oldArgs }()

				singlechecker.Main(cmd2.Analyzer)

			}})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
