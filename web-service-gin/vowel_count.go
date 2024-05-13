package main

func CountVowels(word []string) map[string]map[string]int {
	total_slice := map[string]int{}
	for _, value := range word {
		total := 0
		for _, v := range value {
			switch string(v) {
			case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
				total++
			}
		}
		total_slice[value] = total
	}

	response := map[string]map[string]int{"total_vowels": total_slice}
	return response
}
