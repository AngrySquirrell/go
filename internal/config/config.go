package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(configPath string) ([]LogConfig, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("fichier de configuration non trouvé: %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier de configuration: %w", err)
	}

	var configs []LogConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing du JSON: %w", err)
	}

	if len(configs) == 0 {
		return nil, fmt.Errorf("aucune configuration de log trouvée dans le fichier")
	}

	for i, config := range configs {
		if config.ID == "" {
			return nil, fmt.Errorf("configuration %d: l'ID ne peut pas être vide", i)
		}
		if config.Path == "" {
			return nil, fmt.Errorf("configuration %d: le chemin ne peut pas être vide", i)
		}
		if config.Type == "" {
			return nil, fmt.Errorf("configuration %d: le type ne peut pas être vide", i)
		}
	}

	return configs, nil
}
