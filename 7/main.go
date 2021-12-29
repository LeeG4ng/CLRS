package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	n, k             int               // 任务数 机器数
	taskTime         []float64         // taskTime[i]为任务i的持续时间
	bestStartTime    []float64         // bestStartTime[i]为最优情况任务i的开始时间
	bestMachine      []int             // bestMachine[i]为最优情况任务i分配的机器
	bestEndTime      = math.MaxFloat64 // bestEndTime 为最佳调度的结束时间
	currentStartTime []float64         // currentStartTime[i]为当前任务i开始时间
	currentMachine   []int             // currentMachine[i]为当前任务i分配的机器
	currentSumTime   []float64         // currentSumTime[m]为当前机器m的运行总时间
)

func BackTrack(i int) { // 任务i取值为0,1,...,n-1
	if i == n {
		endTime := finishTime()
		if endTime < bestEndTime {
			bestEndTime = endTime
			copy(bestStartTime, currentStartTime)
			copy(bestMachine, currentMachine)
		}
	} else {
		for m := 0; m < k; m++ {
			currentMachine[i] = m                   // 机器m分配给任务i
			currentStartTime[i] = currentSumTime[m] // 任务i的开始时间为当前机器m的总运行时间
			currentSumTime[m] += taskTime[i]        // 机器m的运行时间加上任务i的持续时间
			if finishTime() < bestEndTime {         // 剪枝
				BackTrack(i + 1)
			}
			currentSumTime[m] -= taskTime[i] // 机器m的运行时间减去任务i的持续时间
		}
	}
}

func finishTime() (max float64) {
	for _, time := range currentSumTime {
		max = math.Max(max, time)
	}
	return max
}

func main() {
	readFile()
	BackTrack(0)

	fmt.Printf("总耗时：%.1f\n调度方案：\n", bestEndTime)
	for j := 0; j < k; j++ {
		fmt.Printf("机器%d（\n", j)
		for i := 0; i < n; i++ {
			if bestMachine[i] == j {
				fmt.Printf("\t任务%d：%.1f\n", i, bestStartTime[i])
			}
		}
		fmt.Printf("）\n")
	}
}

func readFile() {
	file, err := os.Open("7/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	line := fileScanner.Text()
	fmt.Sscanf(line, "%d %d", &n, &k)

	// 根据n、k为全局变量分配空间
	taskTime = make([]float64, n)
	bestStartTime = make([]float64, n)
	bestMachine = make([]int, n)
	currentStartTime = make([]float64, n)
	currentMachine = make([]int, n)
	currentSumTime = make([]float64, k)

	fileScanner.Scan()
	line = fileScanner.Text()
	times := strings.Split(line, " ")
	for i, time := range times {
		taskTime[i], _ = strconv.ParseFloat(time, 64)
	}
	file.Close()
}
