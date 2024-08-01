package compare

import "cmp"

type Func[T any] func(item1, item2 T) bool

func Reversed[T any](f Func[T]) Func[T] {
	return func(item1, item2 T) bool {
		return f(item2, item1)
	}
}

func Equal[T comparable](item1, item2 T) bool {
	return item1 == item2
}

func Less[T cmp.Ordered](item1, item2 T) bool {
	return item1 < item2
}

func Greater[T cmp.Ordered](item1, item2 T) bool {
	return item1 > item2
}

func LessOrEqual[T cmp.Ordered](item1, item2 T) bool {
	return item1 <= item2
}

func GreaterOrEqual[T cmp.Ordered](item1, item2 T) bool {
	return item1 >= item2
}
