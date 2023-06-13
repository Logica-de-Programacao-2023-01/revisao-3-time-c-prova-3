package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums); j++ {

			if j != i && nums[j]+nums[i] == target {

				return []int{i, j}

			}

		}

	}

	return nil

}
