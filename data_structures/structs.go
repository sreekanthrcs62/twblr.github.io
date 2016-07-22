package data_structures

// Linked List Implementation

type Node struct{
	data int
}

type NodeList struct{
	nodeList [] Node
}

func (list *NodeList) Add(elem int) {
	node := Node{data:elem}
	list.nodeList = append(list.nodeList, node)
}

func (list *NodeList) Search(elem int) *Node {

	for _, node := range list.nodeList{
		if node.data == elem {
			return &node
		}
	}
	return nil
}
