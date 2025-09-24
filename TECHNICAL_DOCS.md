# Documentation Technique - GoLog Analyzer

## 📋 Table des matières

- [Architecture générale](#architecture-générale)
- [Packages et modules](#packages-et-modules)
- [Gestion des erreurs](#gestion-des-erreurs)
- [Concurrence](#concurrence)
- [Formats de données](#formats-de-données)
- [Points d'extension](#points-dextension)

## 🏗️ Architecture générale

### Principe de conception

L'application suit les principes de Clean Architecture et les conventions Go :

- **Séparation des responsabilités** : Chaque package a une responsabilité claire
- **Inversion de dépendance** : Les packages internes ne dépendent pas des détails d'implémentation
- **Modularité** : Code réutilisable et testable

### Flux d'exécution

1. **main.go** → Point d'entrée, délègue à Cobra
2. **cmd/root.go** → Configuration générale CLI
3. **cmd/analyze.go** → Logique de la commande analyze
4. **internal/config** → Chargement de la configuration
5. **internal/analyzer** → Analyse parallèle des logs
6. **internal/reporter** → Export des résultats

## 📦 Packages et modules

### `internal/config`

**Responsabilité :** Gestion de la configuration JSON

```go
type LogConfig struct {
    ID   string `json:"id"`   // Identifiant unique du log
    Path string `json:"path"` // Chemin vers le fichier
    Type string `json:"type"` // Type de log (métadonnée)
}
```

**Fonctions principales :**

- `LoadConfig(configPath string) ([]LogConfig, error)` : Charge et valide la configuration

**Validation :**

- Existence du fichier de configuration
- Format JSON valide
- Champs requis présents et non vides

### `internal/analyzer`

**Responsabilité :** Moteur d'analyse et gestion des erreurs

```go
type AnalysisResult struct {
    LogID        string `json:"log_id"`
    FilePath     string `json:"file_path"`
    Status       string `json:"status"`        // "OK" ou "FAILED"
    Message      string `json:"message"`
    ErrorDetails string `json:"error_details"`
}
```

**Erreurs personnalisées :**

```go
type FileNotFoundError struct {
    Path string
}

type ParseError struct {
    Path string
    Err  error
}
```

**Fonctions principales :**

- `AnalyzeLogs(configs []config.LogConfig) []AnalysisResult` : Analyse parallèle
- `analyzeLog(logConfig config.LogConfig) AnalysisResult` : Analyse d'un seul log
- `DisplayResults(results []AnalysisResult)` : Affichage console

### `internal/reporter`

**Responsabilité :** Export et sauvegarde des rapports

**Fonctions principales :**

- `ExportToJSON(results []analyzer.AnalysisResult, outputPath string) error`
- `ExportToJSONWithTimestamp(results []analyzer.AnalysisResult, basePath string) (string, error)`

**Fonctionnalités :**

- Création automatique des dossiers parent
- Format JSON indenté pour lisibilité
- Support des horodatages (format AAMMJJ)

### `cmd/`

**Responsabilité :** Interface utilisateur CLI avec Cobra

**Structure :**

- `root.go` : Commande racine et configuration globale
- `analyze.go` : Commande d'analyse avec drapeaux spécifiques

## 🚨 Gestion des erreurs

### Types d'erreurs

1. **Erreurs système** : Fichiers introuvables, permissions
2. **Erreurs de format** : JSON malformé, champs manquants
3. **Erreurs de parsing** : Contenu de log non analysable (simulé)

### Stratégie de gestion

```go
// Vérification avec types d'erreurs spécifiques
if os.IsNotExist(err) {
    return FileNotFoundError{Path: logConfig.Path}
}

// Wrapping d'erreurs pour contexte
return fmt.Errorf("erreur lors de la lecture du fichier: %w", err)
```

### Messages utilisateur

- **Console** : Messages colorés avec emojis pour clarté
- **JSON** : Détails techniques complets pour débogage
- **Codes de sortie** : Conformes aux conventions Unix

## ⚡ Concurrence

### Architecture concurrente

```go
func AnalyzeLogs(configs []config.LogConfig) []AnalysisResult {
    results := make([]AnalysisResult, len(configs))
    var wg sync.WaitGroup

    // Canal pour collecte sécurisée
    resultChan := make(chan struct {
        result AnalysisResult
        index  int
    }, len(configs))

    // Une goroutine par log
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

    // Synchronisation et collecte
    go func() {
        wg.Wait()
        close(resultChan)
    }()

    for res := range resultChan {
        results[res.index] = res.result
    }

    return results
}
```

### Sécurité des données

- **Channels** : Communication thread-safe entre goroutines
- **WaitGroup** : Synchronisation propre des tâches
- **Index preservation** : Maintien de l'ordre des résultats

### Performance

- **Parallélisme** : Analyse simultanée de tous les logs
- **Simulation réaliste** : Délais aléatoires (50-200ms)
- **Gestion mémoire** : Channels avec buffer pour éviter les blocages

## 📄 Formats de données

### Configuration d'entrée (JSON)

```json
[
  {
    "id": "unique-identifier",
    "path": "/path/to/logfile.log",
    "type": "log-category"
  }
]
```

### Rapport de sortie (JSON)

```json
[
  {
    "log_id": "unique-identifier",
    "file_path": "/path/to/logfile.log",
    "status": "OK|FAILED",
    "message": "Human readable message",
    "error_details": "Technical error details"
  }
]
```

### Validation des données

- **Types stricts** : Validation à l'import et export
- **Champs requis** : Vérification de présence
- **Format cohérent** : Structure identique entrée/sortie

## 🔧 Points d'extension

### Nouveaux types de logs

```go
// Ajouter dans analyzer.go
func analyzeSpecificLogType(logConfig config.LogConfig) AnalysisResult {
    switch logConfig.Type {
    case "nginx-access":
        return analyzeNginxLog(logConfig)
    case "mysql-error":
        return analyzeMySQLLog(logConfig)
    // ... autres types
    default:
        return analyzeGenericLog(logConfig)
    }
}
```

### Nouvelles commandes CLI

```go
// Nouveau fichier cmd/validate.go
var validateCmd = &cobra.Command{
    Use:   "validate",
    Short: "Valider un fichier de configuration",
    Run: func(cmd *cobra.Command, args []string) {
        // Logique de validation
    },
}

func init() {
    rootCmd.AddCommand(validateCmd)
}
```

### Nouveaux formats d'export

```go
// Ajouter dans reporter.go
func ExportToXML(results []analyzer.AnalysisResult, outputPath string) error {
    // Logique d'export XML
}

func ExportToCSV(results []analyzer.AnalysisResult, outputPath string) error {
    // Logique d'export CSV
}
```

### Métriques et monitoring

```go
// Nouveau package internal/metrics
type Metrics struct {
    TotalFiles     int
    SuccessCount   int
    FailureCount   int
    AverageTime    time.Duration
    StartTime      time.Time
    EndTime        time.Time
}
```

## 🧪 Tests recommandés

### Tests unitaires

```go
func TestLoadConfig(t *testing.T) {
    // Test de chargement valide
    // Test de fichier inexistant
    // Test de JSON malformé
    // Test de champs manquants
}

func TestAnalyzeLog(t *testing.T) {
    // Test de fichier existant
    // Test de fichier inexistant
    // Test de permissions
    // Test de simulation d'erreur
}
```

### Tests d'intégration

```go
func TestFullAnalysisWorkflow(t *testing.T) {
    // Test du workflow complet
    // Création de fichiers temporaires
    // Vérification des résultats
    // Nettoyage
}
```

### Tests de performance

```go
func BenchmarkAnalyzeLogs(b *testing.B) {
    // Test de performance avec N logs
    // Mesure du temps d'exécution
    // Analyse de l'utilisation mémoire
}
```

## 🔍 Debugging et monitoring

### Logs de debug

```go
// Ajouter des logs détaillés
log.Printf("Analyse du fichier %s démarrée", logConfig.Path)
log.Printf("Analyse terminée en %v", duration)
```

### Métriques de performance

- Temps d'exécution par log
- Utilisation mémoire
- Nombre de goroutines actives
- Débit de traitement (logs/seconde)

---

Cette documentation technique devrait permettre à tout développeur de comprendre, maintenir et étendre l'application GoLog Analyzer.
