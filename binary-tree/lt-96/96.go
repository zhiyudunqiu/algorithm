package lt96

func numTrees(n int) int {
	var memo = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, n+1)
	}

	return count(1, n, memo)
}

func count(low, hight int, memo [][]int) int {
	if low >= hight {
		return 1
	}

	if memo[low][hight] != 0 {
		return memo[low][hight]
	}

	res := 0
	for i := low; i <= hight; i++ {
		left := count(low, i-1, memo)
		right := count(i+1, hight, memo)

		res += left * right
	}
	memo[low][hight] = res

	return res
}
