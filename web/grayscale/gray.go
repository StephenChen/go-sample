package main

// pass 3/1000
func passed_3_1000() bool {
	userID := 0
	key := hashFunctions(userID) % 1000
	if key <= 2 {
		return true
	}

	return false
}

func hashFunctions(id int) int {
	return 0
}
