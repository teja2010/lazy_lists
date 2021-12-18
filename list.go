package lazy_lists

type Element interface {
	IsElement()
}

type Int int

func (i Int) IsElement() {}

type String string

func (s String) IsElement() {}

// LazyList when evaluated returns an element and rest of the list
type LazyList func() (Element, LazyList)

// Generator builds elements in the list
type Generator func(index int) Element

// NewLazyList builds a lazy list given the generator
func NewLazyList(g Generator) LazyList {
	return func() (Element, LazyList) {
		return lazyList(0, g)
	}
}

func lazyList(i int, g Generator) (Element, LazyList) {
	return g(i), func() (Element, LazyList) {
		return lazyList(i+1, g)
	}
}

// RecurrenceRelation describes how consecutive elements are related
type RecurrenceRelation func(current Element) (next Element)

// NewLazyList2 builds a lazy list given the Recurrence Relation and the first
// element
func NewLazyList2(first Element, r RecurrenceRelation) LazyList {
	return func() (Element, LazyList) {
		return lazyList2(r, first)
	}
}

func lazyList2(r RecurrenceRelation, curr Element) (Element, LazyList) {
	return curr, func() (Element, LazyList) {
		return lazyList2(r, r(curr))
	}
}
