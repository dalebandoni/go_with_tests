package iteration

func Repeat(character string, numTimes int) string {
	var repeated string
	for i := 0; i < numTimes; i++ {
		repeated += character
	}
	return repeated
}
