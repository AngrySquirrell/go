package cmd

import (
	"fmt"
	"log"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse des fichiers de logs selon la configuration fournie",
	Long: `La commande analyze traite plusieurs fichiers de logs en parallèle
selon une configuration JSON fournie. Elle génère un rapport détaillé
des analyses effectuées.

Exemple:
  loganalyzer analyze --config config.json --output report.json
  loganalyzer analyze -c config.json -o report.json`,

	Run: func(cmd *cobra.Command, args []string) {
		// Vérifier que le chemin de config est fourni
		if configPath == "" {
			log.Fatal("Le drapeau --config (-c) est requis")
		}

		fmt.Printf("🔍 Démarrage de l'analyse avec la configuration: %s\n", configPath)

		// Charger la configuration
		logConfigs, err := config.LoadConfig(configPath)
		if err != nil {
			log.Fatalf("Erreur lors du chargement de la configuration: %v", err)
		}

		fmt.Printf("📁 %d fichiers de logs à analyser\n", len(logConfigs))

		results := analyzer.AnalyzeLogs(logConfigs)

		analyzer.DisplayResults(results)

		if outputPath != "" {
			if err := reporter.ExportToJSON(results, outputPath); err != nil {
				log.Fatalf("Erreur lors de l'export: %v", err)
			}
			fmt.Printf("📄 Rapport exporté vers: %s\n", outputPath)
		}

		fmt.Println("✅ Analyse terminée")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier de sortie JSON (optionnel)")

	analyzeCmd.MarkFlagRequired("config")
}
