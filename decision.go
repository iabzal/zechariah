package iskanderzhuma

import "unicode/utf8"

func getMissionDecision(number int) string {
	numbersDesc := map[int]string{
		1: mOne,
		2: mTwo,
		3: mThree,
		4: mFour,
		5: mFive,
		6: mSix,
		7: mSeven,
		8: mEight,
		9: mNine,
	}
	return getValidString(numbersDesc[number])
}

func getMindDecision(number int) string {
	numbersDesc := map[int]string{
		1: mdOne,
		2: mdTwo,
		3: mdThree,
		4: mdFour,
		5: mdFive,
		6: mdSix,
		7: mdSeven,
		8: mdEight,
		9: mdNine,
	}
	return getValidString(numbersDesc[number])
}

func getValidString(word string) string {
	s := word
	if len(word) > 4090 {
		s = word[0:4090] + "..."
	}
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	return s
}
