package list

type Element struct {
	Value string
	next  *Element
	prev  *Element
}

func (e *Element) Next() { return e }

func (e *Element) Prev() { return e }

type List struct {
	currentElement *Element
	head           *Element
}

func (l List) Front() *Element { return nil }

func (l List) AddAt(i int, v string) error {}

func (l List) Get(i int) (*Element, error) { return nil }

func (l List) Len() int { return 0 }
