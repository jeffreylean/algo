package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

// Travel linkedlist and insert values into an array
func TravelAndArray(list *ListNode) []int {
	out := make([]int, 0)
	for list != nil {
		out = append(out, list.Val)
		list = list.Next
	}
	return out
}

// Add value to linkedlist
func AddValue(head *ListNode, value int) *ListNode {
	list := head
	if list == nil {
		return &ListNode{
			Val: value, Next: nil,
		}
	}
	for list.Next != nil {
		temp := list.Next
		list, temp = temp, list
	}
	list.Next = &ListNode{Val: value, Next: nil}
	return head
}

// Get middle value of linkedlist
func GetMid(list *ListNode) *ListNode {
	slow := list
	fast := list.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Merge 2 list and sort them.
func Merge(list1, list2 *ListNode) *ListNode {
	merged := ListNode{}
	curr := &merged

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return merged.Next
}

// Sort single linkedlist using quicksort method, divide and conquer
func SortLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left := head
	right := GetMid(head)

	temp := right.Next
	right.Next = nil
	right = temp

	sorted1 := SortLinkedList(left)
	sorted2 := SortLinkedList(right)
	return Merge(sorted1, sorted2)
}
