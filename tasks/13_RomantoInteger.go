package main

func romanToInt(s string) int {
	wholeSummand := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	summandWithDiminutive := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	sum := 0
	new_s := []rune(s)
	for i := 0; i < len(new_s); {
		if i+2 <= len(new_s) {
			v := new_s[i : i+2]
			el, status := summandWithDiminutive[string(v)]
			if status {
				sum += el
				i += 2
			} else {
				v := new_s[i]
				sum += wholeSummand[string(v)]
				i++
			}
		} else {
			v := new_s[i]
			sum += wholeSummand[string(v)]
			i++
		}
	}
	return sum
}
