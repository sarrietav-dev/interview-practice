package clone

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Neighbors == nil {
		return &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
	}

	created := make(map[int]*Node)

	return clone(created, node)
}

func clone(created map[int]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}

	if old, ok := created[node.Val]; ok {
		return old
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	created[newNode.Val] = newNode

	for _, neighbor := range node.Neighbors {
		if old, ok := created[neighbor.Val]; !ok {
			created[newNode.Val].Neighbors = append(created[newNode.Val].Neighbors, clone(created, neighbor))
		} else {
			created[newNode.Val].Neighbors = append(created[newNode.Val].Neighbors, old)
		}
	}

	return created[newNode.Val]
}
