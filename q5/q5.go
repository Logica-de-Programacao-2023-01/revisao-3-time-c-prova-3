package q5

import "strings"

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {

	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, " ", "")

	if len(s) == 0 {

		return true

	}

	reversedString := ""

	for i := len(s) - 1; i >= 0; i-- {

		reversedString += string(s[i])

	}

	if reversedString == s {

		return true

	} else {

		return false

	}

}
