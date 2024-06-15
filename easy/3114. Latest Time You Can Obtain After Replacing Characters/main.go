package main

func findLatestTime(s string) string {
	if len(s) != 5 {
		return ""
	}

	runeStr := []rune(s)

	if runeStr[0] == '?' && runeStr[1] == '?' {
		runeStr[0], runeStr[1] = '1', '1'
	} else {
		if runeStr[0] == '?' {
			if runeStr[1] <= '1' {
				runeStr[0] = '1'
			} else {
				runeStr[0] = '0'
			}
		}

		if runeStr[1] == '?' {
			if runeStr[0] == '1' {
				runeStr[1] = '1'
			} else {
				runeStr[1] = '9'
			}
		}
	}

	if runeStr[4] == '?' {
		runeStr[4] = '9'
	}
	if runeStr[3] == '?' {
		runeStr[3] = '5'
	}

	return string(runeStr)
}

func main() {

}
