package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd es el comando principal de tu CLI.
var rootCmd = &cobra.Command{
	Use:   "perfilizer",
	Short: "Herramienta CLI para gestionar tu $PROFILE.ps1",
	Run: func(cmd *cobra.Command, args []string) {
		// Si ejecutas `perfilizer` sin subcomando,
		// muestra la ayuda:
		cmd.Help()
	},
}

// Execute arranca la ejecución de rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
