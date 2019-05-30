package problems

func Solution(n uint32) uint32 {
	return binaryGap(n)
}

////////////////////////////////////////////////////////////////////////////////
// Time complexity is O(Log2(N))
//
// As we produce a binary representation of decimal unsigned integer in reverse order we can omit
// analyzing zero bits from the tail of the given number.
//
// Next we travers the "array" of bits in the reverse order and remember every next
// zero bits series' length in order to compare it with any next one.
//
// When closing "true" bit is detected we stop calculating length of the current "false" bits sequence
// and compare it with current maximum length has been detected to the moment. The longest wins.
//
// At the end we need to compare lengths again as in this case the division portion is the "closing" bit.
// Considering reverse order of the analyze it always equals to "true" bit.
//
////////////////////////////////////////////////////////////////////////////////
func binaryGap(n uint32) uint32 {
	// divide by 2 to effectively shift right till the least significant digit is 1 ie. an odd number
	for n >= 1 && n%2 == 0 {
		n /= 2
	}

	var ls uint32 = 0 //current longest sequence of zeros
	var maxLs uint32 = 0 //max longest sequence of zeros
	
	//shift right and record / store the longest sequence (of zeros)
	for n, r := divide(n); n >= 1; n, r = divide(n) {
		if r == 0 {
			ls += 1 //increment local sequence of zeros
		} else {
			maxLs = max(maxLs, ls) //compare to identify largest sequence
			ls = 0
		}
	}

	return max(maxLs, ls)
}

// find max of 2 numbers
func max(x, y uint32) uint32 {
	if x < y {
		return y
	}
	return x
}

// shift right by dividing by 2 and return result and remainder
func divide(n uint32) (uint32, uint32) {
	return n / 2, n % 2
}
