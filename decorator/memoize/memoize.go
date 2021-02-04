package memoize

import "time"

type f func(int) int
type memo map[int] int

func add_to_memo(x int, y int, m memo) {
	m[x] = y
}

func Memoize(fun f) f {
	m := make(memo)
	return func(x int) int {
		if val, present := m[x]; present {
			time.Sleep(10 * time.Second)
			return val
		}
		val := fun(x)
		add_to_memo(x, val, m)
		return val
	}
}
