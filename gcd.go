package numbertheory

// The timeless GCD algorithm by Euclid
func EuclidGCD(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// The vastly more efficient Binary GCD Algorithm by Stein
func SteinGCD(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	var shift uint
	for shift = 0; (a|b)&1 == 0; shift += 1 {
		a = a >> 1
		b = b >> 1
	}
	for a&1 == 0 {
		a = a >> 1
	}
	for {
		for b&1 == 0 {
			b = b >> 1
		}
		if a > b {
			a, b = b, a
		}
		b = b - a
		if b == 0 {
			break
		}
	}
	return a << shift
}
