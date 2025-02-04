package games

import (
	"learngo/utils"
	"time"
)

var Glider = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func LiveGame(matrix [][]int, timeout, iterations int) {
	n := len(matrix)

	// Печатаем начальное состояние
	utils.PrintMatrix(matrix)

	// Основной цикл поколений
	for i := 0; i < iterations; i++ {
		// Создаём новую матрицу для обновления
		newMatrix := make([][]int, n)
		for x := range newMatrix {
			newMatrix[x] = make([]int, n)
		}

		// Обновляем состояния
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				newMatrix[x][y] = calculateState(x, y, matrix)
			}
		}

		// Заменяем старую матрицу на новую
		matrix = newMatrix

		// Ждём перед отображением следующего поколения
		time.Sleep(time.Duration(timeout) * time.Second)
		utils.PrintMatrix(matrix)
	}
}

func calculateState(x, y int, matrix [][]int) int {
	neighbours := calculateNeighbours(x, y, matrix)
	return calculateNewState(matrix[x][y], neighbours)
}

func calculateNeighbours(x, y int, matrix [][]int) int {
	n := len(matrix)
	neighbours := 0

	// Подсчёт количества живых соседей
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue // Пропускаем саму клетку
			}

			// Обработка границ как условно бесконечных
			ni := (i + n) % n
			nj := (j + n) % n

			if matrix[ni][nj] == 1 {
				neighbours++
			}
		}
	}

	return neighbours
}

func calculateNewState(currentState, neighbours int) int {
	// Определение нового состояния
	if currentState == 0 && neighbours == 3 {
		return 1 // Клетка оживает
	}
	if currentState == 1 && (neighbours < 2 || neighbours > 3) {
		return 0 // Клетка умирает
	}
	return currentState // Состояние остаётся неизменным
}
