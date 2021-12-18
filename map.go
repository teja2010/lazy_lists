package lazy_lists

// Function to map an element into another
// e.g. strconv.Itoa converts Int to String
type Map func(i Element) Element

func (l LazyList) Fmap(m Map) LazyList {
	return func() (Element, LazyList) {
		x, xs := l()
		return m(x), xs.Fmap(m)
	}
}
