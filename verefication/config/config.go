package config

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func Load(logger *zap.SugaredLogger) Settings {
	var settings Settings
	env, exists := os.LookupEnv("env")

	if exists {
		filePath := fmt.Sprintf("./.config/%s.json", env)
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		if err := json.Unmarshal(fileBytes, &settings); err != nil {
			logger.Errorf("Unmarshalling error: %s", err)
			os.Exit(1)
		}
	} else {
		logger.Error("env is not exists")
		os.Exit(1)
	}

	return settings
}
