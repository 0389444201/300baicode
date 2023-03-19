package main

import "fmt"

// Sắp xếp mọto linked list bằng phương pháp mergeSort
// Định nghĩa kiểu dữ liệu ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Hàm merge 2 danh sách đã sắp xếp lại với nhau
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// Nếu một trong hai danh sách rỗng, trả về danh sách còn lại
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// Chọn giá trị nhỏ hơn giữa hai danh sách, đệ quy gọi hàm merge cho phần còn lại
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}

// Hàm sắp xếp danh sách liên kết đơn bằng merge sort
func mergeSort(head *ListNode) *ListNode {
	// Nếu danh sách có ít hơn hoặc bằng một phần tử, trả về danh sách ban đầu
	if head == nil || head.Next == nil {
		return head
	}

	// Chia danh sách thành hai phần
	var prev *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	// Đệ quy gọi merge sort cho cả hai nửa danh sách
	left := mergeSort(head)
	right := mergeSort(slow)

	// Hợp nhất hai danh sách đã sắp xếp lại với nhau thành một danh sách mới đã được sắp xếp
	return merge(left, right)
}

func bai3() {
	// Tạo danh sách liên kết đơn
	node1 := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 1, Next: nil}
	node4 := &ListNode{Val: 5, Next: nil}
	node5 := &ListNode{Val: 3, Next: nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	// In danh sách ban đầu
	fmt.Println("Danh sách ban đầu:")
	current := node1
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()

	// Sắp xếp danh sách bằng merge sort
	sortedList := mergeSort(node1)

	// In danh sách đã được sắp xếp
	fmt.Println("Danh sách đã được sắp xếp:")
	current = sortedList
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
