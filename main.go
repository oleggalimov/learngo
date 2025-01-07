package main

import (
	"flag"
	"learngo/games"
	"learngo/utils"
	"time"
)

func main() {

	// Парсим флаги из командной строки
	var timeout int
	var iterations int
	flag.IntVar(&timeout, "timeout", 1, "Таймаут между итерациями")
	flag.IntVar(&iterations, "iterations", 50, "Таймаут между итерациями")
	flag.Parse()

	// simplePrint(matrix)
	games.LiveGame(games.Glider, timeout, iterations)

}

func simplePrint(matrix [][]int) {
	matrixLen := len(matrix)
	dobleMatrixLen := matrixLen * 2
	utils.PrintMatrix(matrix)

	for i := 0; i < dobleMatrixLen; i++ {
		time.Sleep(time.Second)
		if i < matrixLen {
			matrix[i][i] = 1 // Заполняем главную диагональ
		} else {
			row := dobleMatrixLen - i - 1 // Расчет строки для побочной диагонали
			col := i - matrixLen          // Расчет столбца для побочной диагонали
			matrix[row][col] = 1
		}
		utils.PrintMatrix(matrix)
	}
}
