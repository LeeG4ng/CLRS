package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type charNode struct {
	char        rune
	freq        int
	left, right *charNode
}

// implement interface
func (n charNode) key() float64 {
	return float64(n.freq)
}
func (n charNode) setKey(key float64) {
	n.freq = int(key)
}

func huffman(charMap map[rune]int) *charNode {
	q := NewQueue()
	for char, freq := range charMap {
		node := &charNode{char: char, freq: freq}
		q.Insert(node)
	}
	for i := 1; i <= len(charMap)-1; i++ { // n-1次操作
		x := q.ExtractMin().(*charNode)
		y := q.ExtractMin().(*charNode)
		z := &charNode{
			char:  0,
			freq:  x.freq + y.freq,
			left:  x,
			right: y,
		}
		q.Insert(z)
	}
	return q.ExtractMin().(*charNode)
}

func huffmanCode(node *charNode, code string, codeMap map[rune]string) {
	if node.left != nil { // 当前结点非叶子结点
		huffmanCode(node.left, code+"0", codeMap)
		huffmanCode(node.right, code+"1", codeMap)
	} else {
		codeMap[node.char] = code
		fmt.Println(string(node.char), code)
	}
}

func encode(str string, codeMap map[rune]string) int {
	encodeStr := ""
	for _, char := range str {
		encodeStr += codeMap[char]
	}

	file, err := os.Create("6/encode.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufWriter := bufio.NewWriter(file)
	bufWriter.WriteString(encodeStr)
	bufWriter.Flush()
	return len(encodeStr)
}

func main() {
	str := readString()

	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
	}
	charCount := len(charMap)
	strLen := len(str)
	fixedCodeBits := int(math.Ceil(math.Log2(float64(charCount))))
	fixedCodeLen := fixedCodeBits * strLen
	//fmt.Println(charCount, fixedCodeBits, fixedCodeLen)

	tree := huffman(charMap)
	codeMap := make(map[rune]string)
	huffmanCode(tree, "", codeMap)
	huffmanCodeLen := encode(str, codeMap)
	compression := float64(huffmanCodeLen) / float64(fixedCodeLen)
	fmt.Println("Fixed:", fixedCodeLen, "Huffman:", huffmanCodeLen, "Compression:", compression)
}

func readString() string {
	file, err := os.Open("6/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	line := fileScanner.Text()
	return line
}
