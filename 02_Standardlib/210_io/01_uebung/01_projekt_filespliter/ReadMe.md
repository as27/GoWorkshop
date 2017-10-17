# Projekt Filespliter

Vor vielen vielen Jahren, als Daten noch auf Disketten gespeichert wurden, gab es Programme, welche große Dateien in kleinere Einzeldateien auteilten. 

Diese Programme werden _Filespliter_ genannt. Ziel dieser Übung ist es genau so einen _Filespliter_ in Go zu implementieren. 


## Aufgabenstellung

Aus einer beliebigen Datei (jpg, png, txt, doc, zip, ...) werden meherere kleine Dateien (chunks) erzeugt. 

Als Input Parameter bekommt das Programm eine Datei (als Dateiname), die Größe der Chunks in Bytes und einen Output Folder für die Chunks. 

Die Funktion, welche die Input Datei aufbricht soll als Input einen io.Reader und kein File bekommen. Auf diese Weise kann diese auch leichter getestet werden.
 
Da der io.Reader []byte liest, ist die Aufteilung in die Chunks einfach über ein range zu bewerkstelligen.