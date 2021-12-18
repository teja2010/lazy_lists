package lazy_lists

// A list to generate numbers
var Numbers = NewLazyList(func(idx int) Element {
	return Int(idx)
})

func NumbersFrom(start int) LazyList {
	return NewLazyList2(Int(start),
		func(curr Element) (next Element) {
			num := int(curr.(Int))
			return Int(num + 1)
		},
	)
}
