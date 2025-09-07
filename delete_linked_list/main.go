package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	// Сдвигаем первый указатель на n + 1 шагов вперед
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Двигаем оба указателя до конца списка
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Удаляем n-й элемент с конца
	second.Next = second.Next.Next

	return dummy.Next
}
