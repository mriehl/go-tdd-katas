package fibo

func Fibonacci() func() int {
	var prev, prevprev int
	var initialized bool

	return func() int {
		if !initialized {
			initialized = true
			return 0
		}
		if prev == 0 {
			prev = 1
			return 1
		}
		current := prev + prevprev
		prevprev = prev
		prev = current
		return current
	}
}
