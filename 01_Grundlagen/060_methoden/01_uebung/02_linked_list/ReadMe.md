# Linked List

Wir implementieren eine [Linked List](https://de.wikipedia.org/wiki/Liste_(Datenstruktur)) in Go.


Es gibt dabei folgende Datentypen:

* Data (struct, welches die Daten beinhaltet)
* Node
    * next (Pointer auf das folgende Element)
    * data
* Liste
    * currentNode
    * head

Die Liste sollte folgende Methoden implementieren:

* Add(Data)
* Range
        