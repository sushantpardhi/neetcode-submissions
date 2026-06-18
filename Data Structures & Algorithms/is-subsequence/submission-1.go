func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	s1 := s[0]
	sCount := 0
	for i, _ := range t {
		if t[i] == s1 {
			sCount++
			if sCount == len(s){
				return true
			}
			s1 = s[sCount]
		}
	}

	return false
}
