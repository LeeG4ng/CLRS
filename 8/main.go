package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	WHITE = iota
	GRAY
	BLACK
)

type color int

type VertexNode struct {
	name      string
	firstEdge *EdgeNode
	color     color
}

type EdgeNode struct {
	iVex, jVex   int
	iLink, jLink *EdgeNode
}

type Graph struct {
	vList []VertexNode
	vNum  int
}

func (g *Graph) InsertVertex(v string) {
	g.vList = append(g.vList, VertexNode{v, nil, WHITE})
	g.vNum++
}

func (g *Graph) InsertEdge(v1, v2 string) {
	iVex, jVex := -1, -1
	var iNode, jNode *VertexNode
	for index, v := range g.vList {
		if v1 == v.name {
			iVex = index
			iNode = &g.vList[index]
		}
		if v2 == v.name {
			jVex = index
			jNode = &g.vList[index]
		}
	}

	edge := &EdgeNode{
		iVex: iVex,
		jVex: jVex,
	}

	// 链表插入
	edge.iLink = iNode.firstEdge
	iNode.firstEdge = edge
	edge.jLink = jNode.firstEdge
	jNode.firstEdge = edge
}

func (g *Graph) adj(v int) (list []int) {
	e := g.vList[v].firstEdge
	for e != nil {
		if e.iVex == v {
			list = append(list, e.jVex)
			e = e.iLink
		} else { // e.jVex == v
			list = append(list, e.iVex)
			e = e.jLink
		}
	}
	return
}

func (g *Graph) BFS(s int) {
	g.vList[s].color = GRAY
	fmt.Print(g.vList[s].name)
	q := NewQueue()
	q.Push(s)

	for !q.IsEmpty() {
		u := q.Pop().(int)
		for _, v := range g.adj(u) {
			if g.vList[v].color == WHITE {
				g.vList[v].color = GRAY
				fmt.Print("-", g.vList[v].name)
				q.Push(v)
			}
		}
		g.vList[u].color = BLACK
	}
}

func main() {
	vertexes, pairs := readFile()

	g := new(Graph)
	for _, v := range vertexes {
		g.InsertVertex(v)
	}
	for _, p := range pairs {
		g.InsertEdge(p.v1, p.v2)
	}

	//g.BFS(0)
	for i := 0; i < g.vNum; i++ {
		if g.vList[i].color == WHITE {
			g.BFS(i)
			fmt.Print("\n")
		}
	}
}

type pair struct {
	v1, v2 string
}

func readFile() (vertexes []string, pairs []pair) {
	file, _ := os.Open("8/data.txt")
	fileScanner := bufio.NewScanner(file)

	fileScanner.Scan()
	line := fileScanner.Text()
	vertexes = strings.Split(line, ",")

	for fileScanner.Scan() {
		line = fileScanner.Text()
		p := strings.Split(line, "-")
		pairs = append(pairs, pair{p[0], p[1]})
	}
	file.Close()
	return
}
