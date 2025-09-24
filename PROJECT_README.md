# GoLog Analyzer

## 📝 Description

GoLog Analyzer est un outil CLI développé en Go pour analyser des fichiers de logs provenant de diverses sources (serveurs, applications).

## 📦 Installation

### Installation des dépendances

```bash

# Installer les dépendances
go mod tidy

# Compiler l'application
go build -o loganalyzer.exe
```

### Installation des packages

```bash
# Installer Cobra pour l'interface CLI
go get github.com/spf13/cobra@latest
```

## 🎯 Utilisation

### Commande analyze

La commande principale pour analyser les fichiers de logs :

```bash
# Analyse basique avec configuration
./loganalyzer analyze --config config.json

# Analyse avec export du rapport
./loganalyzer analyze --config config.json --output report.json

# Utilisation des raccourcis
./loganalyzer analyze -c config.json -o report.json
```

### Format du fichier de configuration

Le fichier de configuration doit être au format JSON et contenir un tableau de logs à analyser :

```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  }
]
```

### Format du rapport de sortie

Le rapport généré contient les informations suivantes pour chaque log analysé :

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  }
]
```

## 🏗️ Tree

```
├── main.go                    # Point d'entrée principal
├── config.json                # Fichier de configuration d'exemple
├── go.mod                     #
├── cmd/                       # Commandes CLI
│   ├── root.go                #
│   └── analyze.go             #
├── internal/                  # Packages internes
│   ├── config/                #
│   │   └── config.go          #
│   ├── analyzer/              #
│   │   └── analyzer.go        #
│   └── reporter/              #
│       └── reporter.go        #
└── test_logs/                 # Fichiers de logs de test
    ├── access.log             #
    ├── errors.log             #
    ├── empty.log              #
    └── corrupted.log          #
```

## 🧪 Tests

### Fichiers de test inclus

Le projet inclut plusieurs fichiers de test dans `test_logs/` :

- `access.log` - Log d'accès web (valide)
- `errors.log` - Log d'erreurs application (valide)
- `empty.log` - Fichier vide (valide)
- `corrupted.log` - Fichier corrompu (pour tester la gestion d'erreurs)

### Exécuter les tests

```bash
# Tester l'analyse avec la configuration d'exemple
./loganalyzer analyze -c config.json -o test_report.json

# Tester l'affichage de l'aide
./loganalyzer --help
./loganalyzer analyze --help
```

## 🎁 Features

### Horodatage des exports JSON

Support pour nommer les fichiers de sortie avec une date au format AAMMJJ.

```go
// Utilisation dans le code
timestampedPath, err := reporter.ExportToJSONWithTimestamp(results, "report.json")
// Génère: 240924_163505_report.json
```

## 📋 Exemples d'utilisation

### Analyse simple

```bash
./loganalyzer analyze --config config.json
```

### Analyse avec export

```bash
./loganalyzer analyze --config config.json --output rapports/analyse_$(date +%Y%m%d).json
```

### Aide et documentation

```bash
./loganalyzer --help
./loganalyzer analyze --help
```

## 📄 Licence

Ce projet est développé dans le cadre d'un TP pour EFREI Paris.

Exercice original trouvable [ici](https://github.com/axellelanca/loganizer)

---

**Développé avec ❤️ en Go par Guillaume Maugin**
