# LogAnalyzer - Analyseur de Logs Distribuée

## Fonctionnalités

- **Concurrence :** Traitement parallèle via goroutines et WaitGroups
- **Erreurs personnalisées :** Types d'erreurs spécialisés avec `errors.Is()` et `errors.As()`
- **CLI Cobra :** Interface en ligne de commande avec sous-commandes
- **Import/Export JSON :** Configuration d'entrée et rapports de sortie
- **Architecture modulaire :** Code organisé en packages (`internal/`)


## Structure

```
go_loganizer/
├── main.go
├── cmd/
│   ├── root.go
│   └── analyze.go
├── internal/
│   ├── config/config.go
│   ├── analyzer/
│   │   ├── analyzer.go
│   │   └── errors.go       
│   └── reporter/reporter.go
└── test_logs/
```

## Usage

### Commande analyze

```bash
./loganizer analyze --config <path> [--output <path>]
```

**Options :**
- `-c, --config` : Fichier de configuration JSON (obligatoire)
- `-o, --output` : Fichier de rapport JSON (optionnel)

### Fichier de configuration

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

### Rapport de sortie

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log", 
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "fichier introuvable ou inaccessible: /non/existent/log.log"
  }
]
```

## Architecture technique

### Concurrence
- **Goroutines :** Une par fichier de log
- **WaitGroup :** Synchronisation des goroutines  
- **Channels :** Communication sécurisée des résultats

### Erreurs personnalisées

**Types d'erreurs :**
- `FileNotFoundError` : Fichiers introuvables/inaccessibles
- `ParsingError` : Erreurs de parsing des logs

**Gestion avancée :**
- Utilisation d'`errors.Is()` et `errors.As()`
- Messages d'erreur structurés et clairs

### Simulation d'analyse
- Délai aléatoire : 50-200ms par fichier
- Erreur simulée : 10% de chance d'erreur de parsing
- Vérification d'existence des fichiers

## Exemple d'utilisation

```bash
# Analyse simple
./loganizer analyze -c config.json

# Analyse avec export
./loganizer analyze -c config.json -o rapport.json

# Avec création automatique de répertoires
./loganizer analyze -c config.json -o rapports/2024/rapport.json
```
# Équipe

- GAILLARD Maylis
- DAIJARDIN Enolha
- BRUAIRE Tom
