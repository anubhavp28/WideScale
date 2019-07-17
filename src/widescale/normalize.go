package main

import "strings"

/*
Normalize performs the necessary step of converting related tokens (in our case, words) to a
generic form, e.g. words like 'Elephants', 'elephants', 'elephant' will all be converted to
'elephant'. This should improve the result of our search.
*/
func Normalize(word string) string {
	word = strings.Trim(word, ".,-~?!\"'`;:()<>[]{}\\|/=_+*&^%$#@")
	word = strings.ToLower(word)
	return word
}
