package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

type AnalysisResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

type FileNotFoundError struct {
	Path string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier non trouvé: %s", e.Path)
}

type ParseError struct {
	Path string
	Err  error
}

func (e ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing du fichier %s: %v", e.Path, e.Err)
}

func AnalyzeLogs(configs []config.LogConfig) []AnalysisResult {
	results := make([]AnalysisResult, len(configs))
	var wg sync.WaitGroup
	
	resultChan := make(chan struct {
		result AnalysisResult
		index  int
	}, len(configs))

	for i, logConfig := range configs {
		wg.Add(1)
		go func(idx int, cfg config.LogConfig) {
			defer wg.Done()
			
			result := analyzeLog(cfg)
			resultChan <- struct {
				result AnalysisResult
				index  int
			}{result, idx}
		}(i, logConfig)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for res := range resultChan {
		results[res.index] = res.result
	}

	return results
}

func analyzeLog(logConfig config.LogConfig) AnalysisResult {
	result := AnalysisResult{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
	}

	fileInfo, err := os.Stat(logConfig.Path)
	if err != nil {
		if os.IsNotExist(err) {
			result.Status = "FAILED"
			result.Message = "Fichier introuvable."
			result.ErrorDetails = FileNotFoundError{Path: logConfig.Path}.Error()
			return result
		}
		result.Status = "FAILED"
		result.Message = "Erreur d'accès au fichier."
		result.ErrorDetails = err.Error()
		return result
	}

	if !fileInfo.Mode().IsRegular() {
		result.Status = "FAILED"
		result.Message = "Le chemin ne pointe pas vers un fichier régulier."
		result.ErrorDetails = fmt.Sprintf("%s n'est pas un fichier régulier", logConfig.Path)
		return result
	}

	// Simuler l'analyse avec un délai aléatoire (50-200ms)
	analysisTime := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(analysisTime)

	// Simuler occasionnellement une erreur de parsing (10%)
	if rand.Float32() < 0.1 {
		result.Status = "FAILED"
		result.Message = "Erreur de parsing."
		result.ErrorDetails = ParseError{Path: logConfig.Path, Err: fmt.Errorf("format de log non reconnu")}.Error()
		return result
	}

	result.Status = "OK"
	result.Message = "Analyse terminée avec succès."
	result.ErrorDetails = ""

	return result
}

func DisplayResults(results []AnalysisResult) {
	fmt.Println("\n📊 Résultats d'analyse:")
	fmt.Println("=" + string(make([]rune, 50)) + "=")
	
	successCount := 0
	for _, result := range results {
		if result.Status == "OK" {
			successCount++
		}
		
		statusEmoji := "✅"
		if result.Status == "FAILED" {
			statusEmoji = "❌"
		}
		
		fmt.Printf("%s ID: %s\n", statusEmoji, result.LogID)
		fmt.Printf("   📁 Chemin: %s\n", result.FilePath)
		fmt.Printf("   📊 Statut: %s\n", result.Status)
		fmt.Printf("   💬 Message: %s\n", result.Message)
		
		if result.ErrorDetails != "" {
			fmt.Printf("   🔍 Détails: %s\n", result.ErrorDetails)
		}
		fmt.Println()
	}
	
	fmt.Printf("📈 Résumé: %d/%d analyses réussies\n", successCount, len(results))
}
