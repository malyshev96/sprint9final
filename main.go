package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		panic("Размер должен быть больше 0")
	}
	data := make([]int, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range data{
		data[i] = r.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0{
		return 0
	}
	max := data[0]
	for i := 1; i < len(data); i++{
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) <= CHUNKS{
		return 0
	}

	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup

	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++{
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize

		if i == CHUNKS - 1{
			endIndex = len(data)
		}
		
		wg.Add(1)

		go func(chunkIndex, start, end int){
			defer wg.Done()

			chunkMax := data[start]
			for j := start + 1; j < end; j++{
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}
			maxValues[chunkIndex] = chunkMax
		}(i, startIndex, endIndex)
	}

	wg.Wait()

	max := maxValues[0]
	for i := 1; i < len(maxValues); i++{
		if maxValues[i] > max {
			max = maxValues[i]
		}
	}

	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())
}
