package reader

func ReadAll() {

}

func SplitCSVLine(linea string) []string {
	var result []string = make([]string, 20)
	var current string
	var dentroComillas bool

	index := 0

	for _, char := range linea {
		if char == '"' {
			dentroComillas = !dentroComillas
		} else if char == ',' && !dentroComillas {
			result[index] = current
			index++
			current = ""
		} else {
			current += string(char)
		}
	}

	result[index] = current

	return result
}

func SplitCSVLineForBytes(linea *[]byte) *[][]byte {
	var result [][]byte = make([][]byte, 20)
	var current []byte
	var dentroComillas bool

	index := 0

	for _, char := range *linea {
		if char == '"' {
			dentroComillas = !dentroComillas
		} else if char == ',' && !dentroComillas {
			result[index] = current
			index++
			current = nil
		} else {
			//current += string(char)
			current = append(current, char)
		}
	}

	result[index] = current

	return &result
}
