# Linked List

Wir implementieren eine [Linked List](https://de.wikipedia.org/wiki/Liste_(Datenstruktur)) in Go. Die Liste soll eine definierte API besitzen.

Loop durch die Liste soll wie folgt möglich sein:

```Go
    for e := l.Front(); e != nil; e = e.Next() {
        // do something with e.Value
    }
```
[Quelle golang.org/pkg/container/list](https://golang.org/pkg/container/list/#pkg-index)

Es gibt dabei folgende Datentypen:

* Data (struct, welches die Daten beinhaltet)
* Element
    * next, prev = Pointer zu den jeweiligen Elementen
    * Next()*Element
    * Prev()*Element
* Liste
    * currentNode
    * head

Die Liste sollte folgende Methoden implementieren:

* Front()*Element
* AddAt(int,Data)
* PushFront(Data)
* PushBack(Data)
* Range
