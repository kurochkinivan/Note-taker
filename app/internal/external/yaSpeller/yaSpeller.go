package yaspeller

import (
	"strings"
	"unicode/utf8"

	yandexspeller "github.com/kurochkinivan/Note-taker/pkg/yandexSpeller"
)

func CorrectMistakes(text string) (string, error) {
	mistakes, err := yandexspeller.MakeRequst(text)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	runes := []rune(text)
	i := 0

	for _, mistake := range mistakes {
		sb.WriteString(string(runes[i:mistake.Position]))
		sb.WriteString(mistake.S[0])
		i = mistake.Position + utf8.RuneCountInString(mistake.Word)
	}

	sb.WriteString(string(runes[i:]))

	return sb.String(), nil
}


func correctMistakesRune(text string) (string, error) {
	mistakes, err := yandexspeller.MakeRequst(text)
	if err != nil {
		return "", err
	}

	runes := []rune(text)
	var result []rune
	var i int

	for _, mistake := range mistakes {
		result = append(result, runes[i:mistake.Position]...)
		result = append(result, []rune(mistake.S[0])...)
		i = len([]rune(mistake.Word)) + mistake.Position
	}
	result = append(result, runes[i:]...)

	return string(result), nil
}
