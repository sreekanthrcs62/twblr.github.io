package data_structures

import "testing"

func TestContainsElementsAddedToLinkedList(t *testing.T) {
	list := NodeList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	for i := 1; i <= 3; i++ {
		if list.Search(i) == nil {
			t.Errorf("Should have contained a node with value %d", i)
		}
	}
}
