# Makefile pour GoLog Analyzer

# Variables
BINARY_NAME=loganalyzer
BINARY_WIN=$(BINARY_NAME).exe
MAIN_PACKAGE=.

# Commandes par défaut
.PHONY: help build clean run test deps

# Aide
help:
	@echo "Commandes disponibles:"
	@echo "  build      - Compiler l'application"
	@echo "  clean      - Nettoyer les fichiers générés"
	@echo "  run        - Exécuter l'analyse avec la config par défaut"
	@echo "  test       - Tester l'application"
	@echo "  deps       - Installer les dépendances"
	@echo "  help       - Afficher cette aide"

# Installer les dépendances
deps:
	@echo "Installation des dépendances..."
	go mod tidy
	go get github.com/spf13/cobra@latest

# Compiler l'application
build: deps
	@echo "Compilation de l'application..."
	go build -o $(BINARY_WIN) $(MAIN_PACKAGE)
	@echo "✅ Compilation terminée: $(BINARY_WIN)"

# Nettoyer les fichiers générés
clean:
	@echo "Nettoyage..."
	-del $(BINARY_WIN) 2>nul || true
	-del report.json 2>nul || true
	-del test_report.json 2>nul || true
	@echo "✅ Nettoyage terminé"

# Exécuter l'analyse avec la config par défaut
run: build
	@echo "Exécution de l'analyse..."
	.\$(BINARY_WIN) analyze --config config.json --output report.json

# Tester l'application
test: build
	@echo "Tests de l'application..."
	@echo "1. Test de l'aide:"
	.\$(BINARY_WIN) --help
	@echo ""
	@echo "2. Test de l'analyse:"
	.\$(BINARY_WIN) analyze --config config.json --output test_report.json
	@echo "✅ Tests terminés"

# Installation complète
install: clean build test
	@echo "✅ Installation et tests réussis!"
