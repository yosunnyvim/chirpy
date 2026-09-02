package main

import "strings"

func replaceBadWords(body string ) string{
	words := strings.Split(body, " ")
	for i, word:= range words{
		lower:= strings.ToLower(word)
		if (lower== "kerfuffle" || lower== "sharbert" || lower== "fornax"){
			words[i]="****"	
		}

	}
	return strings.Join(words," ") 
}
