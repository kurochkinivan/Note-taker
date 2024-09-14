package yaspeller

import (
	yandexspeller "github.com/kurochkinivan/Note-taker/pkg/yandexSpeller"
)

func CorrectMistakes(text string) (string, error) {
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
		i = i + len([]rune(mistake.Word)) + (mistake.Position - i)
	}
	result = append(result, runes[i:]...)

	return string(result), nil
}
