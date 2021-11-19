package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Point struct {
	name string
	x, y float64
}

func distance(a, b Point) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(dx*dx + dy*dy)
}

func readFile(points *[]Point) {
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
		*points = append(*points, point)
	}
	fmt.Println("输入：", *points)
	file.Close()
}

func closestPair(points []Point) (dist float64, a, b Point) {
	if len(points) == 1 {
		return math.Inf(1), Point{}, Point{}
	}
	if len(points) == 2 {
		a, b = points[0], points[1]
		return distance(a, b), a, b
	}
	mid := len(points) / 2
	leftDist, leftA, leftB := closestPair(points[:mid])
	rightDist, rightA, rightB := closestPair(points[mid:])
	if leftDist < rightDist {
		dist, a, b = leftDist, leftA, leftB
	} else {
		dist, a, b = rightDist, rightA, rightB
	}
	k := 0 // 将横坐标[mid-d, mid+d]范围内的点移动到最前端
	for i, p := range points {
		if math.Abs(p.x-points[mid].x) <= dist {
			points[k], points[i] = points[i], points[k]
			k++
		}
	}
	// 将前端k个点按y升序排序
	sort.Slice(points[:k], func(i, j int) bool { return points[i].y < points[j].y })
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if points[j].y-points[i].y > dist {
				break
			}
			overlapDist := distance(points[i], points[j])
			if overlapDist < dist {
				dist, a, b = overlapDist, points[i], points[j]
			}
		}
	}
	return
}

func main() {
	var points []Point
	readFile(&points)

	// 按x升序排序
	sort.Slice(points, func(i, j int) bool { return points[i].x < points[j].x })
	fmt.Println("排序：", points)
	distance, a, b := closestPair(points)
	fmt.Println(`距离：`, distance, `最近点对：`, a, b)
	calc(points)
}

// 暴力解
func calc(points []Point) {
	min := math.Inf(1)
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			min = math.Min(min, distance(p, points[j]))
		}
	}
	fmt.Println(`暴力解距离：`, min)
}
