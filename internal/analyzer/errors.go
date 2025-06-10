package analyzer

import "fmt"

type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable ou inaccessible: %s", e.Path)
}

type ParsingError struct {
	Path    string
	Details string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("erreur de parsing du fichier %s: %s", e.Path, e.Details)
}

func NewFileNotFoundError(path string) *FileNotFoundError {
	return &FileNotFoundError{Path: path}
}

func NewParsingError(path, details string) *ParsingError {
	return &ParsingError{Path: path, Details: details}
}
