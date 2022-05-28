package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	numbers = generateNumbers(40)
)

const (
	Slots = 4
)

func BubbleSort(integers []int, sg *sync.WaitGroup, i int, m *sync.Mutex) {
	defer sg.Done()

	m.Lock()
	fmt.Printf("partition %d: ", i)
	fmt.Printf("%v \n", integers)
	m.Unlock()

	n := len(integers)
	for {
		var swapped bool = false
		for i := 1; i <= n-1; i++ {
			if integers[i-1] > integers[i] {
				Swap(integers, i-1)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func Swap(integers []int, i int) {
	temp := integers[i]
	integers[i] = integers[i+1]
	integers[i+1] = temp
}

func main() {
	fmt.Printf("entered %d numbers :)\n", len(numbers))
	fmt.Printf("%v\n\n", numbers)

	if len(numbers)%Slots == 0 {
		m := sync.Mutex{}
		sg := sync.WaitGroup{}
		sg.Add(Slots)

		l := len(numbers) / Slots // 1/4 of the array
		for i := 0; i < Slots; i++ {
			start := i * l
			end := l * (i + 1)
			go BubbleSort(numbers[start:end], &sg, i+1, &m)
		}
		sg.Wait()
		fmt.Printf("\nsorted list:\n%v\n", numbers)

	} else {
		fmt.Printf("Number should be divisible by %d\n", Slots)
	}
}

func generateNumbers(n int) []int {
	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		numbers = append(numbers, rand.Intn(250))
	}
	return numbers
}
