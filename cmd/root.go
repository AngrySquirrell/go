package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil d'analyse de logs distribuée",
	Long: `GoLog Analyzer est un outil CLI développé en Go pour analyser
des fichiers de logs.

Exemples :
  loganalyzer analyze --config config.json --output report.json
  loganalyzer analyze -c config.json -o report.json`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de l'exécution de la commande: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Aide pour toggle")
}
