package main 

import "strings"

func cleanInput(text string) []string {
	slice := strings.Fields(text)
	lowerCaseWords := make([]string, len(slice))
	for idx,item := range slice {
		lowerCaseWords[idx] = strings.ToLower(item)
	}
	return lowerCaseWords
}
