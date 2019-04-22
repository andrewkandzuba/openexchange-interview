package problems

func Solution(n uint32) uint32 {
	return binaryGap(n)
}

////////////////////////////////////////////////////////////////////////////////
// Time complexity is O(Log2(N))
// As we produce binary representation in reverse order we must omit analyzing
// zero bit from the tail of the number.
// Then we travers the "array" in reverse order and remember every next
// zero bits series length in order to compare it with any next one.
////////////////////////////////////////////////////////////////////////////////
func binaryGap(n uint32) uint32 {

	for n >=1 && n % 2 == 0 {
		n /= 2
	}

	var mls uint32 = 0
	var maxMls uint32 = 0

	n, r := divide(n)
	for n >= 1 {
		if r == 0 {
			mls += 1
		} else {
			maxMls = max(maxMls, mls)
			mls = 0
		}
		n, r = divide(n)
	}

	return max(maxMls, mls)
}

func max(x, y uint32) uint32 {
	if x < y {
		return y
	}
	return x
}

func divide(n uint32) (uint32, uint32) {
	return n / 2, n % 2
}
