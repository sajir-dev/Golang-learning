package main

// func checkWildCard(str1 string, str2 string) bool {
// 	runeStr1 := []rune(str1)
// 	runeStr2 := []rune(str2)

// 	switch (len(runeStr1) > 0) && (len(runeStr2) > 0) {
// 	case contains(runeStr1, "*"):
// 		for i:=0; i<len(runeStr1); i++ {
// 			if runeStr1[i] != runeStr2[i] {
// 				if runeStr1[i] === rune("*") {
// 					if (runeStr1[i+1] == nil) {
// 						return true
// 					} else {
// 						if
// 					}
// 				}
// 			}
// 		}

// }

func contains(runeStr []rune) (int, rune) {
	chars := "?*"
	rChar := []rune(chars)
	for i, v := range runeStr {
		if v == rChar[0] {
			return i, rChar[0]
		} else if v == rChar[1] {
			return i, rChar[1]
		}
	}
	return -1, 0
}

func compareStrings(str1 string, str2 string) {
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	place, char := contains(runeStr1)

	if place == -1 {
		place, char = contains(runeStr2)
	}

}
