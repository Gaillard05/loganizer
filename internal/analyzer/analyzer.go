package analyzer

import (
	"math/rand"
	"os"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

func AnalyzeLog(cfg config.LogConfig) config.LogReport {
	_, err := os.Stat(cfg.Path)
	if os.IsNotExist(err) {
		return config.LogReport{
			LogID:        cfg.ID,
			FilePath:     cfg.Path,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: err.Error(),
		}
	}

	time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)

	if rand.Float32() < 0.1 {
		return config.LogReport{
			LogID:        cfg.ID,
			FilePath:     cfg.Path,
			Status:       "FAILED",
			Message:      "Erreur de parsing.",
			ErrorDetails: "erreur simulée de parsing",
		}
	}

	return config.LogReport{
		LogID:        cfg.ID,
		FilePath:     cfg.Path,
		Status:       "OK",
		Message:      "Analyse terminé avec succès",
		ErrorDetails: "",
	}
}
