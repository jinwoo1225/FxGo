package fxgo

type Funcable interface {
	any
}

func Map[T Funcable](iterable []T, f func(T) T) []T {
	var results []T
	for _, it := range iterable {
		results = append(results, f(it))
	}
	return results
}

func Filter[T Funcable](iterable []T, f func(T) bool) []T {
	var results []T
	for _, it := range iterable {
		if f(it) {
			results = append(results, it)
		}
	}
	return results
}

func Reduce[T Funcable](iterable []T, init T, f func(T, T) T) T {
	acc := init
	for _, it := range iterable {
		acc = f(acc, it)
	}
	return acc
}
