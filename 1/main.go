package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func insertionSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low, high, k int) {
	if high-low < k {
		return
	}
	p := partition(arr, low, high)
	quickSort(arr, low, p-1, k)
	quickSort(arr, p+1, high, k)
}

func modifiedQuickSort(arr []int, low, high int) {
	quickSort(arr, low, high, 50)
	insertionSort(arr)
}

func main() {
	// 生成测试数据
	w := 10000
	dataSizes := []int{1*w, 10*w, 100*w, 1000*w}
	testData, testDataCopy := make(map[int][]int), make(map[int][]int)
	rand.Seed(time.Now().UnixNano())
	for _, size := range dataSizes {
		for i := 0; i < size; i++ {
			val := rand.Intn(10000000)
			testData[size] = append(testData[size], val)
			testDataCopy[size] = append(testDataCopy[size], val)
		}
	}
	for _, size := range dataSizes {
		// 测试修改后的快速排序
		start := time.Now()
		modifiedQuickSort(testData[size], 0, size - 1)
		fmt.Println("data size", size/w, "w, quicksort: ", time.Now().Sub(start))
		if !sort.IntsAreSorted(testData[size]) {
			fmt.Println("wrong")
		}
		// 测试库函数内置排序
		start = time.Now()
		sort.Ints(testDataCopy[size])
		fmt.Println("data size", size/w, "w, lib sort: ", time.Now().Sub(start))
	}

}
