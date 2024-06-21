package nextLetterBumper

func nextLetterConvert(w string) string {
	output := make([]int32, len(w))
	for i := 0; i < len(w); i++ {
		n := int32(w[i]) + 1
		if (n == 91) || (n == 123) {
			n = n - 26
		}
		output[i] = n
	}
	return string(output)
}
