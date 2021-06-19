package main

//func LengthOfnoneRepeatSubStr(s string) int {
//	lastOccured := make(map[byte]int)
//	start :=0
//	maxLength := 0
//
//	for i, ch := range []byte(s) {
//		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
//			start = lastOccured[ch] + 1
//		}
//		if i - start + 1 > maxLength {
//			maxLength = i - start + 1
//		}
//		lastOccured[ch] = i
//	}
//	return maxLength
//}

func LengthOfnoneRepeatSubStr(s string) int {
	lastOccured := make(map[rune]int)
	start :=0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}

