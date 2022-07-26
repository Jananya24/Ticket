package common

import (
	"regexp"

	cw "github.com/rmarasigan/bus-ticketing/pkg/cw/logger"
)

// RemoveVowel removes vowel letters from the string passed and
// returns a string without the vowel letters.
func RemoveVowel(str string) (string, error) {
	regex, err := regexp.Compile(`[aeyiou]`)
	if err != nil {
		cw.Error(err, &cw.Logs{Code: "RemoveVowel", Message: "Failed to parse regex expression"})

		return "", err
	}

	return regex.ReplaceAllString(str, ""), nil
}