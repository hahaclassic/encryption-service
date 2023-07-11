package mathfunc

// greatest common divisor
func GreatestCommonDiv(a int, b int) int {

	for a != 0 {
		a, b = b%a, a
	}

	return b
}

// Mod Inverse: a * x % m = 1, search x
// Advanced Euclid algorithm
func FindModInverse(a int, m int) int {

	if GreatestCommonDiv(a, m) != 1 {
		return 0
	}

	u1, u2, u3 := 1, 0, a
	v1, v2, v3 := 0, 1, m
	var q int

	for v3 != 0 {
		q = u3 / v3
		v1, v2, v3, u1, u2, u3 = (u1 - q*v1), (u2 - q*v2), (u3 - q*v3), v1, v2, v3
	}

	return u1 % m
}
