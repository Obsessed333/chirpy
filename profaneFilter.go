package main
import (
	"strings"
)
func ProfaneFilter(s string) string{
	words := strings.Split(s, " ")
	badWords:= map[string]struct{}{
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}
	for i, word := range words{
		if _, ok := badWords[strings.ToLower(word)]; ok {
    	words[i] = "****"
		}
	}
	cleanString := strings.Join(words, " ")
	return cleanString
}