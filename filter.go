package lazy_lists

type Filter func(i Element) bool

func (l LazyList) Filter(f Filter) LazyList {
	return func() (Element, LazyList) {
		xs := l
		for {
			var x Element
			x, xs = xs()
			if x == nil { // end of list
				return x, xs
			}
			if f(x) {
				return x, xs.Filter(f)
			}
		}
	}
}
