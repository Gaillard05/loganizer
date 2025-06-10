package analyzer

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

func AnalyzeLog(cfg config.LogConfig) config.LogReport {
	_, err := os.Stat(cfg.Path)
	if os.IsNotExist(err) {
		customErr := NewFileNotFoundError(cfg.Path)
		return config.LogReport{
			LogID:        cfg.ID,
			FilePath:     cfg.Path,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: customErr.Error(),
		}
	}

	time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)

	if rand.Float32() < 0.1 {
		customErr := NewParsingError(cfg.Path, "format de log invalide détecté")
		return config.LogReport{
			LogID:        cfg.ID,
			FilePath:     cfg.Path,
			Status:       "FAILED",
			Message:      "Erreur de parsing.",
			ErrorDetails: customErr.Error(),
		}
	}

	return config.LogReport{
		LogID:        cfg.ID,
		FilePath:     cfg.Path,
		Status:       "OK",
		Message:      "Analyse terminée avec succès",
		ErrorDetails: "",
	}
}

func ProcessLogWithErrorHandling(cfg config.LogConfig) error {
	_, err := os.Stat(cfg.Path)
	if os.IsNotExist(err) {
		return NewFileNotFoundError(cfg.Path)
	}

	// Simuler une erreur de parsing
	if rand.Float32() < 0.1 {
		return NewParsingError(cfg.Path, "format invalide")
	}

	return nil
}

func HandleCustomErrors(err error) string {
	if err == nil {
		return "Aucune erreur"
	}

	var fileErr *FileNotFoundError
	if errors.As(err, &fileErr) {
		return "Gestion spéciale pour fichier manquant: " + fileErr.Path
	}

	var parseErr *ParsingError
	if errors.As(err, &parseErr) {
		return "Gestion spéciale pour erreur de parsing: " + parseErr.Details
	}

	return "Erreur non gérée: " + err.Error()
}
