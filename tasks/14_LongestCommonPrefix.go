package main

func getIndexSmallestLengthWord(strs []string) int {
	lenShortestWord := len(strs[0])
	idxShortestWord := 0
	for idx, elem := range strs {
		if lenShortestWord > len(elem) {
			idxShortestWord = idx
			lenShortestWord = len(elem)
		}
	}
	return idxShortestWord
}

func longestCommonPrefix(strs []string) string {
	idxSmallestLength := getIndexSmallestLengthWord(strs)
	start := 0
	result := ""
	for end := 1; end <= len(strs[idxSmallestLength]); end++ {
		prefix := strs[idxSmallestLength][start:end]
		contains := true
		for _, elem := range strs {
			elem_prefix := elem[start:end]
			if elem_prefix != prefix {
				contains = false
			}
		}
		if contains {
			result = prefix
		}
	}
	return result
}

func main() {

}
