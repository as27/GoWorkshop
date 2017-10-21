# Projekt Filespliter

Vor vielen vielen Jahren, als Daten noch auf Disketten gespeichert wurden, gab es Programme, welche große Dateien in kleinere Einzeldateien auteilten. 

Diese Programme werden _Filespliter_ genannt. Ziel dieser Übung ist es genau so einen _Filespliter_ in Go zu implementieren. 

## Ziele

Mit diesem Projekt sollst Du verschiedene Fähigkeiten erweben bzw. üben.

* Verwenden der Dokumentation der Standar library
* Verwenden des io.Readers
* Erstellen eigener Tests

## Aufgabenstellung

Aus einer beliebigen Datei (jpg, png, txt, doc, zip, ...) werden meherere kleine Dateien (chunks) erzeugt. 

Als Input Parameter bekommt das Programm eine Datei (als Dateiname), die Größe der Chunks in Bytes und einen Output Folder für die Chunks. Die Eingangswerte können zuerst als Varaiblen deklariert werden. Nachdem der Splitter funktioniert strukturieren wir das Programm um. Dafür wird dann das [flag Packet](https://golang.org/pkg/flag/) zu verwendet.

Die Funktion, welche die Input Datei in kleine Stücke zerlegt soll als Input einen io.Reader bekommen. Diese Umsetzung mit einem io.Reader abstrahiert die Eingabe. 

Die grundlegende Struktur für die _split()_ ist bereits vorgegeben.

Für diese Funktion ist bereits auch schon ein Test vorgegeben.

 
Da der io.Reader []byte liest, ist die Aufteilung in die Chunks einfach über ein range zu bewerkstelligen.