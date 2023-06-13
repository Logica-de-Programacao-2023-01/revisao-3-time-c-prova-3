package bonus

//Você recebe as cabeças de duas listas ligadas ordenadas, list1 e list2.
//
//Una as duas listas em uma única lista ordenada. A lista resultante deve ser construída juntando os nós das duas
//primeiras listas.
//
//Retorne a cabeça da lista ligada mesclada.

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func MergeTwoLists(list1 *LinkedList, list2 *LinkedList) *LinkedList {

	newList := LinkedList{

		Head: nil,
	}

	for list1 != nil && list2 != nil {

		if list1.Head.Value <= list2.Head.Value {

			newList.Head.Next = list1.Head
			list1.Head = list1.Head.Next

		} else {

			newList.Head.Next = list2.Head
			list1.Head = list2.Head.Next

		}

		newList.Head = newList.Head.Next

		if list1 != nil {

			newList.Head.Next = list1.Head

		} else if list2 != nil {
			newList.Head.Next = list2.Head
		}

	}

	return &newList

}
