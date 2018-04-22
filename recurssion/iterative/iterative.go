package iterative

// FiboR - Fibonnacci in a recurssion version
func FiboR(n int) int {
	if n < 2 {
		return n
	}
	return FiboR(n-2) + FiboR(n-1)
}

// FiboI -  Fibonnacci in an iterative version
func FiboI(n int) int {
	var result int

	for i, first, second := 0, 0, 1; i <= n; i, first, second = i+1, first+second, first {
		if i == n {
			result = first
		}
	}
	return result
}

// FiboTail -Fibonnacci in a tail-recursive version
func FiboTail(n int) int {
	return fiboT(n, 0, 1)
}

func fiboT(n, first, second int) int {
	if n == 0 {
		return first
	}
	return fiboT(n-1, second, first+second)
}
