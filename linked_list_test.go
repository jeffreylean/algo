package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTravelAndArray(t *testing.T) {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	assert.ElementsMatch(t, TravelAndArray(&list), []int{1, 2, 3, 4})
}

func TestGetMid(t *testing.T) {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	assert.ElementsMatch(t, TravelAndArray(GetMid(&list)), []int{2, 3, 4})
}

func TestAddValue(t *testing.T) {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	assert.ElementsMatch(t, TravelAndArray(AddValue(&list, 5)), []int{1, 2, 3, 4, 5})
}

func TestMerge(t *testing.T) {
	list1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	list2 := ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}
	assert.ElementsMatch(t, TravelAndArray(Merge(&list1, &list2)), []int{1, 2, 2, 3, 3, 4, 4, 5})
}

func TestSortLinkedList(t *testing.T) {
	list1 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: nil}}}}
	assert.ElementsMatch(t, TravelAndArray(SortLinkedList(&list1)), []int{1, 3, 4, 5})
}
