package main

import "fmt"

type Node struct {
	size   int
	key    [3]int
	first  *Node
	second *Node
	third  *Node
	fourth *Node
	parent *Node
}

func NewNode(k int) *Node {
	return &Node{
		size:   1,
		key:    [3]int{k, 0, 0},
		first:  nil,
		second: nil,
		third:  nil,
		fourth: nil,
		parent: nil,
	}
}

func NewNodeWithChildren(k int, first, second, third, fourth, parent *Node) *Node {
	return &Node{
		size:   1,
		key:    [3]int{k, 0, 0},
		first:  first,
		second: second,
		third:  third,
		fourth: fourth,
		parent: parent,
	}
}

func (n *Node) find(k int) bool {
	for i := 0; i < n.size; i++ {
		if n.key[i] == k {
			return true
		}
	}
	return false
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func sort2(x, y *int) {
	if *x > *y {
		swap(x, y)
	}
}

func sort3(x, y, z *int) {
	if *x > *y {
		swap(x, y)
	}
	if *x > *z {
		swap(x, z)
	}
	if *y > *z {
		swap(y, z)
	}
}

func (n *Node) sort() {
	switch n.size {
	case 1:
		return
	case 2:
		sort2(&n.key[0], &n.key[1])
	case 3:
		sort3(&n.key[0], &n.key[1], &n.key[2])
	}
}

func (n *Node) insertToNode(k int) {
	n.key[n.size] = k
	n.size++
	n.sort()
}

func (n *Node) removeFromNode(k int) {
	if n.size >= 1 && n.key[0] == k {
		n.key[0] = n.key[1]
		n.key[1] = n.key[2]
		n.size--
	} else if n.size == 2 && n.key[1] == k {
		n.key[1] = n.key[2]
		n.size--
	}
}

func (n *Node) becomeNode2(k int, first, second *Node) {
	n.key[0] = k
	n.first = first
	n.second = second
	n.third = nil
	n.fourth = nil
	n.parent = nil
	n.size = 1
}

func (n *Node) isLeaf() bool {
	return n.first == nil && n.second == nil && n.third == nil
}

func split(item *Node) *Node {
	if item.size < 3 {
		return item
	}

	x := &Node{key: [3]int{item.key[0]}, first: item.first, second: item.second, parent: item.parent}
	y := &Node{key: [3]int{item.key[2]}, first: item.third, second: item.fourth, parent: item.parent}

	if x.first != nil {
		x.first.parent = x
	}
	if x.second != nil {
		x.second.parent = x
	}
	if y.first != nil {
		y.first.parent = y
	}
	if y.second != nil {
		y.second.parent = y
	}

	if item.parent != nil {
		item.parent.insertToNode(item.key[1])

		if item.parent.first == item {
			item.parent.first = nil
		} else if item.parent.second == item {
			item.parent.second = nil
		} else if item.parent.third == item {
			item.parent.third = nil
		}

		if item.parent.first == nil {
			item.parent.fourth = item.parent.third
			item.parent.third = item.parent.second
			item.parent.second = y
			item.parent.first = x
		} else if item.parent.second == nil {
			item.parent.fourth = item.parent.third
			item.parent.third = y
			item.parent.second = x
		} else {
			item.parent.fourth = y
			item.parent.third = x
		}

		tmp := item.parent
		// Here we would typically free the item, but Go has garbage collection
		return tmp
	} else {
		x.parent = item
		y.parent = item
		item.becomeNode2(item.key[1], x, y)
		return item
	}
}

func insert(p *Node, k int) *Node {
	if p == nil {
		return NewNode(k) // Create the first 2-3 node (root) if the tree is empty
	}

	if p.isLeaf() {
		p.insertToNode(k) // Insert the key into the leaf node
	} else if k <= p.key[0] {
		p.first = insert(p.first, k) // Recursively insert into the first child
	} else if (p.size == 1) || (p.size == 2 && k <= p.key[1]) {
		p.second = insert(p.second, k) // Insert into the second child
	} else {
		p.third = insert(p.third, k) // Insert into the third child
	}

	return split(p) // Split the node if necessary and return the new root
}

func search(p *Node, k int) *Node {
	if p == nil {
		return nil
	}

	if p.find(k) {
		return p
	} else if k < p.key[0] {
		return search(p.first, k)
	} else if (p.size == 2 && k < p.key[1]) || p.size == 1 {
		return search(p.second, k)
	} else if p.size == 2 {
		return search(p.third, k)
	}
	return nil
}

func searchMin(p *Node) *Node {
	if p == nil {
		return nil
	}
	if p.first == nil {
		return p
	}
	return searchMin(p.first)
}

func remove(p *Node, k int) *Node {
	item := search(p, k) // Ищем узел, где находится ключ k

	if item == nil {
		return p // Если ключ не найден, возвращаем дерево без изменений
	}

	var min *Node
	if item.key[0] == k {
		min = searchMin(item.second) // Ищем эквивалентный ключ в правом поддереве
	} else {
		min = searchMin(item.third) // Ищем эквивалентный ключ в левом поддереве
	}

	if min != nil { // Меняем ключи местами
		var z *int
		if k == item.key[0] {
			z = &item.key[0]
		} else {
			z = &item.key[1]
		}
		swap(z, &min.key[0]) // Обмен значениями ключей
		item = min           // Перемещаем указатель на лист, т.к. min - всегда лист
	}

	item.removeFromNode(k) // Удаляем требуемый ключ из листа
	return fix(item)       // Вызываем функцию для восстановления свойств дерева
}

func fix(leaf *Node) *Node {
	if leaf.size == 0 && leaf.parent == nil { // Случай 0, когда удаляем единственный ключ в дереве
		// Освобождаем память (в Go это происходит автоматически при сборке мусора)
		return nil
	}

	if leaf.size != 0 { // Случай 1, когда вершина, в которой удалили ключ, имела два ключа
		if leaf.parent != nil {
			return fix(leaf.parent)
		}
		return leaf
	}

	parent := leaf.parent

	// Случай 2, когда достаточно перераспределить ключи в дереве
	if parent.first.size == 2 && parent.second.size == 2 && parent.size == 2 {
		leaf = redistribute(leaf)
	} else if parent.size == 2 && parent.third.size == 2 {
		leaf = redistribute(leaf)
	} else { // Случай 3, когда нужно произвести склеивание
		leaf = merge(leaf)
	}

	return fix(leaf)
}

func redistribute(leaf *Node) *Node {
	parent := leaf.parent
	first := parent.first
	second := parent.second
	third := parent.third

	if parent.size == 2 && first.size < 2 && second.size < 2 && third.size < 2 {
		if first == leaf {
			parent.first = parent.second
			parent.second = parent.third
			parent.third = nil

			parent.first.insertToNode(parent.key[0])
			parent.first.third = parent.first.second
			parent.first.second = parent.first.first

			if leaf.first != nil {
				parent.first.first = leaf.first
			} else if leaf.second != nil {
				parent.first.first = leaf.second
			}

			if parent.first.first != nil {
				parent.first.first.parent = parent.first
			}

			parent.removeFromNode(parent.key[0])
			// Free the first node if necessary (Go handles memory differently)
		} else if second == leaf {
			first.insertToNode(parent.key[0])
			parent.removeFromNode(parent.key[0])
			if leaf.first != nil {
				first.third = leaf.first
			} else if leaf.second != nil {
				first.third = leaf.second
			}

			if first.third != nil {
				first.third.parent = first
			}

			parent.second = parent.third
			parent.third = nil

			// Free the second node if necessary
		} else if third == leaf {
			second.insertToNode(parent.key[1])
			parent.third = nil
			parent.removeFromNode(parent.key[1])
			if leaf.first != nil {
				second.third = leaf.first
			} else if leaf.second != nil {
				second.third = leaf.second
			}

			if second.third != nil {
				second.third.parent = second
			}

			// Free the third node if necessary
		}
	} else if parent.size == 2 && (first.size == 2 || second.size == 2 || third.size == 2) {
		if third == leaf {
			if leaf.first != nil {
				leaf.second = leaf.first
				leaf.first = nil
			}

			leaf.insertToNode(parent.key[1])
			if second.size == 2 {
				parent.key[1] = second.key[1]
				second.removeFromNode(second.key[1])
				leaf.first = second.third
				second.third = nil
				if leaf.first != nil {
					leaf.first.parent = leaf
				}
			} else if first.size == 2 {
				parent.key[1] = second.key[0]
				leaf.first = second.second
				second.second = second.first
				if leaf.first != nil {
					leaf.first.parent = leaf
				}
			}
		}

		// Additional logic for handling redistribution can be added here
	}

	return leaf
}

func merge(leaf *Node) *Node {
	parent := leaf.parent

	if parent.first == leaf {
		parent.second.insertToNode(parent.key[0])
		parent.second.third = parent.second.second
		parent.second.second = parent.second.first

		if leaf.first != nil {
			parent.second.first = leaf.first
		} else if leaf.second != nil {
			parent.second.first = leaf.second
		}

		if parent.second.first != nil {
			parent.second.first.parent = parent.second
		}

		parent.removeFromNode(parent.key[0])
		parent.first = nil
		parent.first = nil
	} else if parent.second == leaf {
		parent.first.insertToNode(parent.key[0])

		if leaf.first != nil {
			parent.first.third = leaf.first
		} else if leaf.second != nil {
			parent.first.third = leaf.second
		}

		if parent.first.third != nil {
			parent.first.third.parent = parent.first
		}

		parent.removeFromNode(parent.key[0])
		parent.second = nil
		parent.second = nil
	}

	if parent.parent == nil {
		var tmp *Node
		if parent.first != nil {
			tmp = parent.first
		} else {
			tmp = parent.second
		}
		tmp.parent = nil
		parent = nil
		return tmp
	}

	return parent
}

func main() {
	tree := NewNode(50)
	insert(tree, 5)
	insert(tree, 4)
	insert(tree, 3)
	insert(tree, 15)
	insert(tree, 25)

	fmt.Println(search(tree, 25))
	remove(tree, 25)
	fmt.Println(search(tree, 25))
	fmt.Println(search(tree, 15))
}
