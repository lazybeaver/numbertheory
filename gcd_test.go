package numbertheory

import "testing"

func checkEuclidGCD(t *testing.T, a int, b int, expected int) {
	actual := EuclidGCD(a, b)
	if actual != expected {
		t.Errorf("EuclidGCD(%d, %d) = %d | Actual=%d", a, b, expected, actual)
	}
}

func checkSteinGCD(t *testing.T, a int, b int, expected int) {
	actual := SteinGCD(a, b)
	if actual != expected {
		t.Errorf("SteinGCD(%d, %d) = %d | Actual=%d", a, b, expected, actual)
	}
}

func checkAll(t *testing.T, a int, b int, expected int) {
	// Euclid
	checkEuclidGCD(t, a, b, expected)
	checkEuclidGCD(t, b, a, expected)
	// Stein
	checkSteinGCD(t, a, b, expected)
	checkSteinGCD(t, b, a, expected)
}

func TestAlgorithms(t *testing.T) {
	checkAll(t, 1, 0, 1)
	checkAll(t, 1, 1, 1)
	checkAll(t, 1, 2, 1)
	checkAll(t, 2, 4, 2)
	checkAll(t, 10, 3, 1)
	checkAll(t, 15, 10, 5)
	checkAll(t, 17, 101, 1)
	checkAll(t, 17, 101, 1)
	checkAll(t, 321451443876, 12555473728888, 4)
}

func BenchmarkEuclid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = EuclidGCD(2147483647, 1099511627791)
		_ = EuclidGCD(310135939356302189, 671845585623995875)
	}
}

func BenchmarkStein(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SteinGCD(2147483647, 1099511627791)
		_ = SteinGCD(310135939356302189, 671845585623995875)
	}
}
