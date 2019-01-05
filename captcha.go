package captcha

func getOperation(input byte) string {
	switch input {
		case '1':
			return "+"
		case '2':
			return "-"
		case '3':
			return "x"
	}
	return ""
}

var mapper = map[byte] string {
	'0': "zero",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
}

func mapIntToStr(input byte) string {
	return mapper[input]
}

func getNumberText(input byte, returnWord bool) string {
	if returnWord {
		return mapIntToStr(input)
	}

	return string(input)
}

func Generate(input string) string {
	var operand = getOperation(input[2])
	wordFirst :=  input[0] == '2'
	return getNumberText(input[1], wordFirst) + " " + operand + " " + getNumberText(input[3], !wordFirst)
}
