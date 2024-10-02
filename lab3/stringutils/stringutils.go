package stringutils

func Str_Reverse(str string) string {
	runeOrig := []rune(str)
	for i := 0; i < len(runeOrig)/2; i++ {
		temp := runeOrig[i]
		runeOrig[i] = runeOrig[len(runeOrig)-1-i]
		runeOrig[len(runeOrig)-1-i] = temp
	}
	return string(runeOrig)
}
