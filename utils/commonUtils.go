package utils

import "fmt"

func PrintMatrix(matrix [][]int) {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			line := matrix[i]
			var symbol string
			if line[j] == 0 {
				symbol = ". "
			} else {
				symbol = "# "
			}
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}
