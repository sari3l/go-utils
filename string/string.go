package string

import "strings"

const (
	Whitespace     = " \t\n\r\v\f"
	AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
	AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AsciiLetters   = AsciiLowercase + AsciiUppercase
	Digits         = "0123456789"
	Hexdigits      = Digits + "abcdef" + "ABCDEF"
	Octdigits      = "01234567"
	Punctuation    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	Printable      = Digits + AsciiLetters + Punctuation + Whitespace
)

func CapWords(str string, sep string) string {
	if sep == "" {
		sep = " "
	}
	var result []string
	for _, s := range strings.Split(str, sep) {
		if len(s) == 0 {
			continue
		}
		result = append(result, strings.ToUpper(s[0:1])+strings.ToLower(s[1:]))
	}
	return strings.Join(result, sep)
}
