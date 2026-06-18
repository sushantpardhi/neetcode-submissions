func isAnagram(s string, t string) bool {
	// If length are not same then the strings cannot be Anagram
	if len(s) != len(t) {
		return false
	}

	// Creating hash map for first string
	freq1 := make(map[rune]int)
	for _, ch := range s{
		freq1[ch]++
	}

	// Creating hashmap for second string
	freq2 := make(map[rune]int)
	for _, ch := range t{
		freq2[ch]++
	}

	// Comparing the character count for specific ch in both hashmap
	for ch, count := range freq1{
		if freq2[ch] != count{
			return false
		}
	}

	return true
}