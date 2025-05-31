package games

import (
	"bufio"
	"log"
	"os"
)

func readMapFromFile(path string) ([][]int, error) {
	if path == "" {
		path = "data/map.data"
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %s", path)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	battleMap := make([][]int, 0, 10)

	for i := 0; i < 10 && scanner.Scan(); i++ {
		line := scanner.Text()
		runes := []rune(line)     // Преобразуем в руны для работы с символами
		row := make([]int, 0, 19) // Создаём строку матрицы
		for j := 0; j < len(runes) && j < 19; j++ {
			char := runes[j]
			if char == ' ' {
				continue
			}
			symbol := 0
			if char == 'x' {
				symbol = 1
			}
			row = append(row, symbol)
		}
		battleMap = append(battleMap, row)
	}
	return battleMap, nil
}

func validatePositions(battleMap [][]int) bool {
	m := len(battleMap)
	n := len(battleMap[0])
	log.Printf("Проверяем расстановку на поле %dx%d", m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !validatePosition(battleMap, m, n, i, j) {
				return false
			}
		}
	}

	return true
}

func buildBattleShips(battleMap [][]int) bool {

}

func validatePosition(battleMap [][]int, rows, columns, i, j int) bool {
	if battleMap[i][j] != 1 {
		return true
	}
	for row := i - 1; row <= i+1; row = row + 2 {
		if row < 0 || row >= rows {
			continue
		}
		for col := j - 1; col <= j+1; col = col + 2 {
			if col >= columns || col < 0 {
				continue
			}
			if battleMap[row][col] == 1 {
				log.Fatalf("В позиции %d:%d найдена ошибка", row+1, col+1)
				return false
			}
		}
	}

	return true

}

func BattleShipGame() {
	battleMap, _ := readMapFromFile("")
	validatePositions(battleMap)
	log.Print("Проверка завершена")
	// 	grid := [][] int{
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	// 		{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	}

}

// func BattleShip() {
// 	players := list.New()

// 	players.PushBack(1)
// 	players.PushBack(2)
// 	players.PushBack(3)

// 	printList(players)
// 	first := players.Front()
// 	players.Remove(first)
// 	printList(players)

// }

// func printList(l *list.List) {
// 	for element := l.Front(); element != nil; element = element.Next() {
// 		fmt.Printf("%v ", element.Value)
// 	}
// 	fmt.Println()
// }
