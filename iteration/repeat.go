package iteration

func Repeat(n int, character string) (sequence string) {
	for i := 1; i <= n; i++ {
		sequence += character
	}

	return
}