package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
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

func generateHuffmanTree(arrayNodes []Node) (tree Node) {
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

func createEncodeString(data []byte, codes map[uint16]string) (compressed string) {
	for _, vl := range data {
		compressed += codes[uint16(vl)]
	}
	return compressed
}

func createEncodedFile(fileName string, compress string, frequency []uint16) {
	var bt2 uint8
	var bitsBuffer = 0

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer out.Close()

	bytesCreated := 0
	lastBits := len(compress) % 8

	var bytesToWrite []byte

	buf := new(bytes.Buffer)
	for _, vl := range frequency {
		err := binary.Write(buf, binary.LittleEndian, uint16(vl))
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	bytesToWrite = buf.Bytes()

	for _, vl := range compress {
		if vl == '0' {
			bt2 = bt2 << 1
		}
		if vl == '1' {
			bt2 = bt2<<1 + 1
		}
		bitsBuffer++
		if bitsBuffer == 8 {
			bytesCreated++
			bytesToWrite = append(bytesToWrite, bt2)
			bitsBuffer = 0
			bt2 = 0
		}
	}
	for i := 0; i < (8 - lastBits); i++ {
		bt2 = bt2 << 1
	}
	bytesCreated++
	bytesToWrite = append(bytesToWrite, bt2)

	err = ioutil.WriteFile(fileName, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}

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

func huffmanTree(frequency []uint16) (root Node) {

	arrayNodes := initializeNodes(frequency)

	sort.Slice(arrayNodes, func(i, j int) bool {
		if arrayNodes[i].Freq == arrayNodes[j].Freq {
			return arrayNodes[i].Letter < arrayNodes[j].Letter
		} else {
			return arrayNodes[i].Freq < arrayNodes[j].Freq
		}
	})

	root = generateHuffmanTree(arrayNodes)
	return root
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
