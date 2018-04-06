package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Letter uint16
	Freq   uint16
	Esq    *Node
	Dir    *Node
}

type Code struct {
	Code byte
}

func (n Node) initializeNode(l uint16, f uint16, e *Node, d *Node) {
	n = Node{l, f, e, d}
}

func initializeNodes(frequency []uint16) (arrayNodes []Node) {
	for i := 0; i < 256; i++ {
		if frequency[i] > 0 {
			arrayNodes = append(arrayNodes, Node{Letter: uint16(i), Freq: frequency[i], Esq: nil, Dir: nil})
		}
	}
	return arrayNodes
}

func huffmanTree(frequency []uint16) (root Node) {
	arrayNodes := initializeNodes(frequency)
	sort.Slice(arrayNodes, func(i, j int) bool {
		if arrayNodes[i].Freq == arrayNodes[j].Freq {
			return arrayNodes[i].Letter < arrayNodes[j].Letter
		} else {
			return arrayNodes[i].Freq < arrayNodes[j].Freq
		}
	})
	root = generateNewRoot(arrayNodes)
	return root
}

func generateNewRoot(arrayNodes []Node) (tree Node) {
	var n Node
	for len(arrayNodes) > 1 {
		n = createNode(&arrayNodes[0], &arrayNodes[1])
		arrayNodes[1] = n
		arrayNodes = append(arrayNodes[:0], arrayNodes[1:]...)
		sort.Slice(arrayNodes, func(i, j int) bool {
			if arrayNodes[i].Freq == arrayNodes[j].Freq {
				return arrayNodes[i].Letter < arrayNodes[j].Letter
			} else {
				return arrayNodes[i].Freq < arrayNodes[j].Freq
			}
		})
	}
	return arrayNodes[0]
}

func createNode(n1 *Node, n2 *Node) (n Node) {
	n = Node{Letter: 257, Freq: n1.Freq + n2.Freq, Esq: &Node{Letter: n1.Letter, Freq: n1.Freq, Esq: n1.Esq, Dir: n1.Dir}, Dir: &Node{Letter: n2.Letter, Freq: n2.Freq, Esq: n2.Esq, Dir: n2.Dir}}
	return n
}

func generateCodes(tree Node, cds map[uint16]string) {

	var walkTree func(n *Node, code string, cds map[uint16]string)

	walkTree = func(n *Node, code string, cds map[uint16]string) {
		if n.Esq == nil {
			cds[n.Letter] = code
			return
		}
		code += "0"
		walkTree(n.Esq, code, cds)
		code = code[:len(code)-1]
		code += "1"
		walkTree(n.Dir, code, cds)
		code = code[:len(code)-1]
	}
	var code string
	walkTree(&tree, code, cds)
}

func showPreOrder(tree *Node) {
	fmt.Printf("root: %v \n", tree)

	var preOrder func(tree *Node)
	preOrder = func(tree *Node) {

		if (*tree).Esq == nil {
			fmt.Printf("%v \n", *tree)
			return
		}

		if (*tree).Esq != nil {
			preOrder(tree.Esq)
		}
		if (*tree).Dir != nil {
			preOrder(tree.Dir)
		}

	}
	preOrder(tree.Esq)
	preOrder(tree.Dir)
}

func countRemainingLeafs(frequency []uint16) (count int) {
	count = 0
	for _, vl := range frequency {
		if vl > 0 {
			count++
		}
	}
	return count
}
