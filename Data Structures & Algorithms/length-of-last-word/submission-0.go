func lengthOfLastWord(s string) int {
	cnt := 0
	lastCnt := 0

	for _, ch := range s {
		if ch >= 65 && ch <= 122 {
			cnt++
			lastCnt = cnt
		} else {
			cnt = 0
		}
	}

	if cnt > lastCnt {
		return cnt
	} else {
		return lastCnt
	}
}
