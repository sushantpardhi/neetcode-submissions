func countSeniors(details []string) int {
	cnt := 0
	for _, detail := range details { 
		age, _ := strconv.Atoi(detail[11:13])
		if  age > 60 {
			cnt++
		}
	}

	return cnt
}