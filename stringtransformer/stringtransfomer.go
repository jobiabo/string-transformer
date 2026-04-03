package stringtransformer

import (
	"strings"
)

func Transform(action string, words []string) string {
	switch action {
	case "upper":
		return UpperCases(words)
	case "lower":
		return LowerCases(words)
	case "cap":
		return CapFirstLetter(words)
	case "snake":
		return strings.ToLower(Snake(words))
	case "reverse":
		return Reverse(words)
	case "title":
		return Title(words)
	default:
		return "Invalid command"
	}
}

func UpperCases(s []string) string {
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] != "upper" {
			result = append(result, s[i])
		}

	}

	finalresult := strings.Join(result, " ")

	return strings.ToUpper(finalresult)

}

func LowerCases(s []string) string {
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] != "lower" {
			result = append(result, s[i])
		}

	}

	finalresult := strings.Join(result, " ")

	return strings.ToLower(finalresult)

}

func CapFirstLetter(s []string) string {
	// var ArticleChange []string
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToUpper(string(string(s[i][0]))) + strings.ToLower(string(string(s[i][1:])))

	}

	return strings.Join(s, " ")
}

func Snake(s []string) string {
	// b := strings.Fields(s)
	added := strings.Join(s, "_")

	return added

}

func Reverse(s []string) string {
	b := strings.Join(s, " ")
	var result string

	for i := len(b) - 1; i >= 0; i-- {
		result += string(b[i])
	}

	return result
}

func Title(s []string) string {
	for i := 0; i < len(s); i++ {
		if s[0] == "the" || s[0] == "an" || s[0] == "on" || s[0] == "a" || s[0] == "and" || s[0] == "but" || s[0] == "or" || s[0] == "for" || s[0] == "nor" ||
			(s[0]) == "at" || s[0] == "to" || s[0] == "by" || s[0] == "in" || s[0] == "of" || s[0] == "as" || s[0] == "up" || s[0] == "is" || s[0] == "it" || s[0] == "am" || s[0] == "i" {
			s[i] = strings.ToUpper(string(s[0][0])) + s[0][1:]

		} else if i > 0 && (s[i] != "the" && s[i] != "an" && s[i] != "on" && s[i] != "a" && s[i] != "and" && s[i] != "but" && s[i] != "or" && s[i] != "for" && s[i] != "nor" &&
			s[i] != "at" && s[i] != "to" && s[i] != "by" && s[i] != "in" && s[i] != "of" && s[i] != "as" && s[i] != "up" && s[i] != "is" && s[i] != "it" && s[i] != "am" && s[i] != "i") {
			s[i] = strings.ToUpper(string(s[i][0])) + s[i][1:]
		} else {
			s[0] = strings.ToUpper(string(s[0][0])) + s[0][1:]
		}
	}

	return strings.Join(s, " ")

}
