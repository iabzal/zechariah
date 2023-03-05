package zechariah

import "unicode/utf8"

func getYearDecision(number int, gender string) string {
	var numbersDesc map[int]string
	if gender == man {
		numbersDesc = map[int]string{
			1: manYearOne,
			2: manYearTwo,
			3: manYearThree,
			4: manYearFour,
			5: manYearFive,
			6: manYearSix,
			7: manYearSeven,
			8: manYearEight,
			9: manYearNine,
		}
	} else if gender == woman {
		numbersDesc = map[int]string{
			1: womanYearOne,
			2: womanYearTwo,
			3: womanYearThree,
			4: womanYearFour,
			5: womanYearFive,
			6: womanYearSix,
			7: womanYearSeven,
			8: womanYearEight,
			9: womanYearNine,
		}
	}
	return getValidString(numbersDesc[number])
}

func getMindDecision(number int, gender string) string {
	var numbersDesc map[int]string
	if gender == man {
		numbersDesc = map[int]string{
			1: manMindOne,
			2: manMindTwo,
			3: manMindThree,
			4: manMindFour,
			5: manMindFive,
			6: manMindSix,
			7: manMindSeven,
			8: manMindEight,
			9: manMindNine,
		}
	} else if gender == woman {
		numbersDesc = map[int]string{
			1: womanMindOne,
			2: womanMindTwo,
			3: womanMindThree,
			4: womanMindFour,
			5: womanMindFive,
			6: womanMindSix,
			7: womanMindSeven,
			8: womanMindEight,
			9: womanMindNine,
		}
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
