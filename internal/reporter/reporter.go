package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
)

func ExportToJSON(results []analyzer.AnalysisResult, outputPath string) error {

	// Encoder les résultats en JSON avec indentation pour la lisibilité
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON: %w", err)
	}

	// Écrire le fichier
	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier: %w", err)
	}

	return nil
}

func ExportToJSONWithTimestamp(results []analyzer.AnalysisResult, basePath string) (string, error) {
	dir := filepath.Dir(basePath)
	ext := filepath.Ext(basePath)
	nameWithoutExt := basePath[:len(basePath)-len(ext)]
	
	now := time.Now()
	dateStr := now.Format("060102_150405")
	
	timestampedPath := fmt.Sprintf("%s_%s%s", dateStr, filepath.Base(nameWithoutExt), ext)
	fullPath := filepath.Join(dir, timestampedPath)

	err := ExportToJSON(results, fullPath)
	return fullPath, err
}
