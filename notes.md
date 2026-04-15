# Rclone Plugin Development

## Original code - Nicht ändern

```bash
/Users/ludal/src/github.com/rclone/rclone
/Users/ludal/src/github.com/rclone/rclone/backend
```

## Plugin Struktur

Eigentliche Plugins liegen im `plugins/` Verzeichnis. Jedes Plugin hat sein eigenes Unterverzeichnis mit einer `main.go` Datei.
Beispiel: `plugins/example/main.go`

## How to build

Um Plugins für rclone zu schreiben, musst du folgende Schritte befolgen:

1. **Erstelle ein neues Go-Paket**: Schreibe dein Backend-Paket so, als ob es im-tree wäre, aber setze den Paketnamen auf "main".

2. **Baue das Plugin**: Baue das Plugin mit dem Befehl:
   ```bash
   go build -buildmode=plugin -o librcloneplugin_NAME.so
   ```
   wobei NAME dem fs.RegInfo.Name des Plugins entspricht.

3. **Lade das Plugin**: Wenn die Umgebungsvariable `$RCLONE_PLUGIN_PATH` gesetzt ist, werden alle Go-Plugins in diesem Verzeichnis mit dem Namen `librcloneplugin_NAME.so` geladen.

## Entwicklung mit lokalem rclone

Falls du mit dem lokalen rclone-Quellcode arbeitest, führe vor dem Bauen folgende Schritte aus:

```bash
cd plugins/example
go mod edit -replace github.com/rclone/rclone=../../../../rclone/rclone
go mod tidy
```

Diese Schritte sind in der Dokumentation des `plugin`-Pakets detailliert beschrieben.
