# Typen

Go ist eine Sprache mit einer [strengen Typisierung](https://de.wikipedia.org/wiki/Starke_Typisierung). Das bedeutet, dass der Typ einer Variablen im Code mit angegeben werden muss. Bei Sprachen mit einer schwachen Typisierung wird der Typ der Variablen zur Laufzeit des Programmes erst ermittelt. 



Anders ausgedrückt, wenn einmal einer Variablen ein Typ zugewiesen wird, so kann dieser auch nicht mehr verändert werden.

In JavaScript (oder auch Python) wird dies nicht so streng gehandhabt. Dies führt in der Praxis immer wieder zu sehr schwer auffindbaren Fehlern. Das ist auch der Grund, warum sich in den letzten Jahren immer mehr [TypeScript](https://de.wikipedia.org/wiki/TypeScript) durchsetzt. TypeScript ist praktisch JavaScript mit Typen.

Der Vorteil der starken Typisierung ist 

* eine erhöhte Lesbarkeit, da bei der Variablendeklaration zusätlich auch der Typ mitgegeben wird
* stabilerer Code, da der Compiler ein Mischen von Typen nicht zulässt. 


Sollte in Go einer Variablen mit dem Typ `string` (also Text) ein `integer` (also Ganzzahl) zugewiesen werden, so lässt sich das Programm gar nicht kompilieren. Der Compiler meldet hier frühzeitig einen Fehler (https://play.golang.org/p/9OdsHARtZE).

```Go
var a string
a = 1
```
```
cannot use 1 (type int) as type string in assignment
```

Eine gute Übersicht über alle Basistypen bietet die Go Tour: https://go-tour-de.appspot.com/basics/11

In der [Dokumentation der Packetes `builtin`](https://golang.org/pkg/builtin/) werden die einzelnen Typen auch genau beschrieben.

Für die erste Verwendung im Workshop beschränken wir uns auf die Typen:

* bool (kann true oder false sein)
* string (representiert einen Text)
* int (ist eine Ganzzahl)
* byte (ein Byte, wird z.B. verwendet, wenn eine Datein eingelesen wird)

## Deklaration

Die Deklaration von Typen in Go sieht wie folgt aus:


```Go
var myString string
myString = "mein Text"
var myInteger int
myInteger = 4
// Inklusive Initialisierung
var anotherInteger int = 4
```

[Beispiele aus der Go Tour](https://go-tour-de.appspot.com/basics/8)

Es gibt auch die Möglichkeit einer kurzen Variablendeklaration. Hierbei ist die Lesbarkeit des Codes ein wenig schlechter.

```Go
myInteger := 4
```

## Eigene Typen

Ein weiterer wichtiger Bestandteil ist die Definition von eigenen Typen.
