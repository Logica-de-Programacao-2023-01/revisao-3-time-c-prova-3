package q2

import "strings"

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	maior := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(maior) > 0 && !strings.HasPrefix(strs[i], maior) {
			maior = maior[:len(maior)-1]
		}

		if maior == "" {
			break
		}
	}

	return maior

}
