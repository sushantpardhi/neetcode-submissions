func hasDuplicate(nums []int) bool {
    freq := make(map[int]int)

	for _, num := range nums{
		freq[num]++

		if freq[num] == 2{
			return true
		}
	}

	return false
}
