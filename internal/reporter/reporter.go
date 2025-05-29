package reporter

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/axellelanca/go_loganizer/internal/config"
)

func ExportJSON(reports []config.LogReport, path string) error {

	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(reports, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
