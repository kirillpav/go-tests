package iteration



func Repeat(count int, character string) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}