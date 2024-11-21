package main

import (
	"fmt"
	"time"
)

func sumArray(arr []int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

func sumArrayToChan(arr []int, ch chan int) {
	ch <- sumArray(arr)
}

func sumArrayByChan(arr []int) int {
	ch := make(chan int)

	go sumArrayToChan(arr[:len(arr)/2], ch)
	go sumArrayToChan(arr[len(arr)/2:], ch)

	sum1 := <-ch
	sum2 := <-ch

	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)

	fmt.Println("Total Sum:", sum1+sum2)

	return sum1 + sum2
}

func main() {
	var arr []int = make([]int, 100)

	for i := 0; i < 100; i++ {
		arr[i] = i
	}

	start := time.Now()

	fmt.Println("Sum:", sumArray(arr))

	fmt.Println("Duration:", time.Since(start))

	startParallel := time.Now()

	sumArrayByChan(arr)

	fmt.Println("Duration:", time.Since(startParallel))
}
