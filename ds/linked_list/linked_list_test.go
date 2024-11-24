package linked_list

import (
	"reflect"
	"slices"
	"testing"
)

func TestToSlice(t *testing.T) {
	ll := ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: nil}}}

	actual := ll.ToSlice()
	expected := []int{1, 2, 3}

	if !slices.Equal(actual, expected) {
		t.Fatalf("actual %v did not equal expected %v", actual, expected)
	}
}

func TestFromSlice(t *testing.T) {
	actual := FromSlice([]int{1, 2, 3})
	expected := &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: nil}}}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("actual %v did not equal expected %v", actual, expected)
	}
}
