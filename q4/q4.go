package q4

//Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
//elemento.
//
//Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

func SingleNumber(nums []int) int {

	frequencias := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		frequencias[nums[i]]++

	}

	single := 0

	for num, fr := range frequencias {

		if fr == 1 {

			single = num

		}

	}

	for i := 0; i < len(nums); i++ {

		if nums[i] == single {

			return i

		}

	}

	return 0

}
