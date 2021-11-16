package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type color string

const (
	RED   = "RED"
	BLACK = "BLACK"
)

type TNode struct {
	color          color
	key            int
	left, right, p *TNode
}

type RBTree struct {
	root, nil *TNode
}

func NewRBTree() *RBTree {
	nilNode := TNode{BLACK, 0, nil, nil, nil}
	return &RBTree{
		root: &nilNode,
		nil:  &nilNode,
	}
}

func (t *RBTree) LeftRotate(node *TNode) {

}

func (t *RBTree) RightRotate(node *TNode) {

}

func (t *RBTree) Insert(node *TNode) {
	y, x := t.nil, t.root
	for x != t.nil { // 循环找到插入的位置，y保持为x的父节点
		y = x
		if node.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	// y为待插入节点的父节点
	node.p = y
	if y == t.nil { //此时t为空树
		t.root = node
	} else if node.key < y.key {
		y.left = node
	} else {
		y.right = node
	}
	// 已经设置过color left right
	t.InsertFixup(node)
}

func (t *RBTree) InsertFixup(node *TNode) {

}

func readFile() []int {
	var (
		count int
		data  []int
	)
	file, err := os.Open("3/insert.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	line := fileScanner.Text()
	fmt.Sscanf(line, "%d", count)
	fileScanner.Scan()
	line = fileScanner.Text()
	elements := strings.Fields(line)
	for _, element := range elements {
		n, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, n)
	}
	fmt.Println("输入：", data)
	file.Close()
	return data
}

func printTree(t *RBTree) {
	q := NewQueue()
	q.Push(t.root)
	for top := q.Pop(); top != nil; top = q.Pop() {
		node := top.(*TNode)
		fmt.Println(node.key, node.color)
		if node.left != t.nil {
			q.Push(node.left)
		}
		if node.right != t.nil {
			q.Push(node.right)
		}
	}
}

func main() {
	data := readFile()
	t := NewRBTree()
	for _, num := range data {
		t.Insert(&TNode{RED, num, t.nil, t.nil, t.nil})
	}
	printTree(t)
}
