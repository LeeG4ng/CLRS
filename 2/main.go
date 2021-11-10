package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	name string
	x, y float64
}

func main() {
	var points []Point
	pointRegexp := regexp.MustCompile(`^([\w]+)\((\d*\.?\d*), (\d*\.?\d*)\)$`)

	file, err := os.Open("2/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		params := pointRegexp.FindStringSubmatch(line)
		x, err := strconv.ParseFloat(params[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseFloat(params[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		point := Point{name: params[1], x: x, y: y}
		points = append(points, point)
	}
	fmt.Println(points)
	file.Close()
}
