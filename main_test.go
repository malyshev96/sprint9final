package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T){
	data := generateRandomElements(10)
	assert.Len(t, data, 10, "Длина слайса не соответствует заданной")
	allZero := false
	for _, num := range data{
		if num != 0{
			break
		}
		allZero = true
	}
	assert.False(t, allZero, "Все элементы нулевые")
	
	assert.Panics(t, func() {generateRandomElements(-2)}, "Паника при отрицательном размере")
	assert.Panics(t, func() {generateRandomElements(0)}, "Паника при нулевом размере")
	assert.Len(t, generateRandomElements(1), 1, "Длина слайса должна быть 1")
}

func TestMaximum(t *testing.T) {
	t.Run("Положительные числа", func(t *testing.T) {
		data := []int{1, 5, 3, 9, 2}
		result := maximum(data)
		assert.Equal(t, 9, result, "Максимум должен быть 9")
	})

	t.Run("Отрицательные числа", func(t *testing.T) {
		data := []int{-5, -1, -10, -3}
		result := maximum(data)
		assert.Equal(t, -1, result, "Максимум должен быть -1")
	})

	t.Run("Любые числа", func(t *testing.T) {
		data := []int{-10, 0, 15, -5, 20}
		result := maximum(data)
		assert.Equal(t, 20, result, "Максимум должен быть 20")
	})

	t.Run("Один элемент", func(t *testing.T) {
		data := []int{42}
		result := maximum(data)
		assert.Equal(t, 42, result, "Максимум должен быть 42")
	})

	t.Run("Одинаковые числа", func(t *testing.T) {
		data := []int{7, 7, 7, 7, 7}
		result := maximum(data)
		assert.Equal(t, 7, result, "Максимум должен быть 7")
	})

	t.Run("Пустой слайс", func(t *testing.T) {
		data := []int{}
		result := maximum(data)
		assert.Equal(t, 0, result, "Должен возвращаться 0")
	})

	t.Run("Максимум в начале слайса", func(t *testing.T) {
		data := []int{100, 50, 75, 25}
		result := maximum(data)
		assert.Equal(t, 100, result, "Максимум должен быть 100")
	})

	t.Run("Максимум в конце слайса", func(t *testing.T) {
		data := []int{10, 20, 30, 40}
		result := maximum(data)
		assert.Equal(t, 40, result, "Максимум должен быть 40")
	})
}