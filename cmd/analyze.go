package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les logs à partir d'un fichier de configuartion JSON",
	Run: func(cmd *cobra.Command, args []string) {
		runAnalyze()
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Fichier de sortie JSON (facultatif)")
	analyzeCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(analyzeCmd)
}

func runAnalyze() {
	logs, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Erreur de chargement : %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	resultsChan := make(chan config.LogReport, len(logs))

	for _, log := range logs {
		wg.Add(1)
		go func(l config.LogConfig) {
			defer wg.Done()
			result := analyzer.AnalyzeLog(l)
			resultsChan <- result
		}(log)
	}

	wg.Wait()
	close(resultsChan)

	var results []config.LogReport
	for r := range resultsChan {
		results = append(results, r)
		fmt.Printf("ID: %s\nChemin: %s\nStatut: %s\nMessage: %s\n",
			r.LogID, r.FilePath, r.Status, r.Message)
		if r.ErrorDetails != "" {
			fmt.Printf("Erreur: %s\n", r.ErrorDetails)
		}
		fmt.Println("-----")
	}

	if outputPath != "" {
		err := reporter.ExportJSON(results, outputPath)
		if err != nil {
			fmt.Printf("Erreur d'export : %v\n", err)
		} else {
			fmt.Printf("Rapport exporté dans %s\n", outputPath)
		}
	}
}
