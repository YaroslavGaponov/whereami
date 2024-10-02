package flags

import (
	"errors"
)

const offset = 0x1F1E6 - 'A'

func GetCountryFlag(code string) (string, error) {
	if len(code) != 2 {
		return "", errors.New("country code is incorrect")
	}
	runes := []rune(code)
	return string(runes[0]+offset) + string(runes[1]+offset), nil
}
