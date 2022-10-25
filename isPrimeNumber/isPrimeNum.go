package isPrimeNumber

func isPrimeNum(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= num / 2; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}