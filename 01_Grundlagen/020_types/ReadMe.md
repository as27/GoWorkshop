# Typen

In Go spielen Typen eine sehr wichtige Rolle. Eine Variable kann dabei immer nur einen bestimmten Typ haben. Wenn einmal einer Variablen ein Typ zugewiesen wird, so kann dieser auch nicht mehr verändert werden.

In Python oder JavaScript wird dies nicht so streng gehandhabt. Diese strenge Typisierung ist am Anfang für manch einen Programmierer ungewohnt. Aber die Verwendung von Typen hat große Vorteile bezüglich lesbarkeit und prüfbarkeit. 

So wird in JavaScript Projekten immer häufiger TypeScript eingesetzt, da hier schwer auffindbare Bugs frühzeitig vermieden werden. D.h. die Typisierung hält auch in diese anderen Sprachen langsam Einzug.

Sollte in Go einer Variablen mit dem Typ `string` (also Text) ein `integer` (also Ganzzahl) zugewiesen werden, so lässt sich das Programm gar nicht kompilieren. Der Compiler meldet hier frühzeitig einen Fehler (https://play.golang.org/p/9OdsHARtZE).

```Go
var a string
a = 1
```
```
cannot use 1 (type int) as type string in assignment
```

