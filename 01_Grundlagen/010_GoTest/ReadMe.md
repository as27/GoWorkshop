# Go Test

Go unterstützt von natur aus automatische Tests. Diese Tests machen wir uns beim Workshop bei den Übungen zu nutze. Innerhalb der Übungen sind Tests vorbereitet. Diese Tests prüfen dann, ob die Aufgaben auch korrekt gelöst wurden. Auf diese Weise kann jeder selber kontrollieren, ob seine Lösung auch korrekt ist.

Das ist auch der Grund, warum vor allen anderen Elementen die Tests gleich als erstes vorgestellt werden.

## Die Testdatei

Eine Testdatei wird immer zu einer Quelldatei erzeugt. Dafür wird dem Dateinamen `_test` angehängt.

Tests für `main.go` liegen immer in der Datei `main_test.go` oder bei `server.go` in `server_test.go`.

Auf diese Weise werden die zugehörigen Testdateien im Verzeichnig gleich bei der Go Datei angezeigt.

## Test über die Commandozeile

Die Go Tests werden meistens über einen Terminal gestartet. Der Befehl hierfür ist `go test`. Damit werden alle tests des aktuellen Verzeichnises getestet.

Möchte man alle Tests eines Projektes ausführen so verwendet man.

```
go test ./...
```

## Test in Visual Studio Code

VS Code unterstützt. 