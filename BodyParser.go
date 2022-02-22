package BodyParser

import (
	"strings"
)

func Parser(s string) map[string]string {
	s = s + `&`
	maps := make(map[string]string)
	chars := []rune(s)
	countList := []int{}
	for i, c := range chars {
		if string(c) == `&` {
			countList = append(countList, i)
		}
	}
	lastIndex := 0
	for _, x := range countList {
		tmpString := s[lastIndex:x]
		lastIndex = x
		var tmpKey string
		var tmpValue string

		tmpChars := []rune(tmpString)
		for index, x := range tmpChars {
			if string(x) == `=` {
				tmpKey = tmpString[:index]
				tmpKey = strings.ReplaceAll(tmpKey, `&`, ``)
				tmpKey = strings.ReplaceAll(tmpKey, `=`, ``)
				tmpValue = tmpString[index:]
				tmpValue = strings.ReplaceAll(tmpValue, `&`, ``)
				tmpValue = strings.ReplaceAll(tmpValue, `=`, ``)
				maps[tmpKey] = tmpValue
			}
		}
	}
	return maps
}
