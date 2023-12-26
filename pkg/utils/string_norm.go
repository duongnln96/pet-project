package utils

import "golang.org/x/text/unicode/norm"

func UnicodeNorm(input string) string {
	return norm.NFC.String(input)
}
