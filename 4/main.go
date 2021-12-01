package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	low, high int
}

type color string

const (
	RED   = "RED"
	BLACK = "BLACK"
)

type TNode struct {
	color          color
	key            int
	left, right, p *TNode

	int interval
	max int
}

type RBTree struct {
	root, nil *TNode
}

func NewRBTree() *RBTree {
	nilNode := TNode{BLACK, 0, nil, nil, nil, interval{0, 0}, 0}
	return &RBTree{
		root: &nilNode,
		nil:  &nilNode,
	}
}

func (t *RBTree) LeftRotate(node *TNode) {
	y := node.right
	node.right = y.left
	y.left.p = node
	y.p = node.p
	if node.p == t.nil {
		t.root = y
	} else if node == node.p.left {
		node.p.left = y
	} else {
		node.p.right = y
	}
	y.left = node
	node.p = y

	// !
	y.max = node.max
	node.max = max(node.int.high, node.left.max, node.right.max)
}

func (t *RBTree) RightRotate(node *TNode) {
	y := node.left
	node.left = y.right
	y.right.p = node
	y.p = node.p
	if node.p == t.nil {
		t.root = y
	} else if node == node.p.left {
		node.p.left = y
	} else {
		node.p.right = y
	}
	y.right = node
	node.p = y

	// !
	y.max = node.max
	node.max = max(node.int.high, node.left.max, node.right.max)
}

func (t *RBTree) Insert(z *TNode) {
	y, x := t.nil, t.root
	for x != t.nil { // 循环找到插入的位置，y保持为x的父节点
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
		y.max = max(y.max, z.max) // !
	}
	// y为待插入节点的父节点
	z.p = y
	if y == t.nil { //此时t为空树
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	// 已经设置过color left right
	t.InsertFixup(z)
}

func (t *RBTree) InsertFixup(z *TNode) {
	for z.p.color == RED {
		if z.p == z.p.p.left { // z.p是个左孩子
			y := z.p.p.right    // y为z的叔节点
			if y.color == RED { // case1：叔节点为红色
				z.p.color = BLACK
				y.color = BLACK   // z的父节点和叔节点改为黑色
				z.p.p.color = RED // z.p.p改为红色
				z = z.p.p         // z指向z.p.p，进入下次循环
			} else {
				if z == z.p.right { // case2：z是一个右孩子
					z = z.p
					t.LeftRotate(z)
				} // case2结束进入case3
				z.p.color = BLACK // case3：z是一个左孩子
				z.p.p.color = RED // 交换p和p.p的颜色
				t.RightRotate(z.p.p)
			}
		} else { // z.p是个右孩子
			y := z.p.p.left     // y为z的叔节点
			if y.color == RED { // case4：叔节点为红色
				z.p.color = BLACK
				y.color = BLACK   // z的父节点和叔节点改为黑色
				z.p.p.color = RED // z.p.p改为红色
				z = z.p.p         // z指向z.p.p，进入下次循环
			} else {
				if z == z.p.left { // case5：z是一个左孩子
					z = z.p
					t.RightRotate(z)
				} // case5结束进入case6
				z.p.color = BLACK // case6：z是一个右孩子
				z.p.p.color = RED // 交换p和p.p的颜色
				t.LeftRotate(z.p.p)
			}
		}
	}
	t.root.color = BLACK
}

func readFile() []interval {
	var (
		count int
		data  []interval
	)
	file, err := os.Open("4/insert.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	line := fileScanner.Text()
	fmt.Sscanf(line, "%d", count)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		elements := strings.Fields(line)
		low, _ := strconv.Atoi(elements[0])
		high, _ := strconv.Atoi(elements[1])
		data = append(data, interval{low, high})
	}
	file.Close()
	fmt.Println("输入：", data)
	return data
}

func printTree(t *RBTree) {
	file, err := os.Create("4/LOT.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufWriter := bufio.NewWriter(file)

	q := NewQueue()
	q.Push(t.root)
	for top := q.Pop(); top != nil; top = q.Pop() {
		node := top.(*TNode)
		if node == t.nil {
			bufWriter.WriteString("nil\n")
			continue
		}
		bufWriter.WriteString(fmt.Sprintf("%d, %s, %d\n", node.int, node.color, node.max))
		q.Push(node.left)
		q.Push(node.right)
	}
	bufWriter.Flush()
}

func (t *RBTree) intervalSearch(x *TNode, i interval) []interval {
	var res []interval
	if overlap(i, x.int) {
		res = append(res, x.int)
	}
	if x.left != t.nil && x.left.max >= i.low {
		res = append(res, t.intervalSearch(x.left, i)...)
	}
	if x.right != t.nil && x.right.max >= i.low && x.int.low <= i.high {
		res = append(res, t.intervalSearch(x.right, i)...)
	}
	return res
}

func main() {
	data := readFile()
	t := NewRBTree()
	for _, interval := range data {
		t.Insert(&TNode{
			color: RED,
			key:   interval.low,
			left:  t.nil,
			right: t.nil,
			p:     t.nil,
			int:   interval,
			max:   interval.high,
		})
	}
	printTree(t)

	for {
		fmt.Print("Input interval:")
		var interval interval
		fmt.Scanf("%d %d", &interval.low, &interval.high)
		res := t.intervalSearch(t.root, interval)
		fmt.Println(res)
	}
}

func max(vals ...int) int {
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func overlap(i1, i2 interval) bool {
	return i1.low <= i2.high && i2.low <= i1.high
}
