package utils

type memoized[P comparable, R interface{}] struct {
	f     func(P) R
	cache map[P]R
}

func (mem *memoized[P, R]) Call(x P) R {
	if v, ok := mem.cache[x]; ok {
		return v
	}
	result := mem.f(x)
	mem.cache[x] = result
	return result
}

func Memoize[P comparable, R interface{}](f func(P) R) *memoized[P, R] {
	return &memoized[P, R]{f: f, cache: make(map[P]R)}
}
