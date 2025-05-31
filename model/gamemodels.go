package model

type Point struct {
	X, Y int
}

type BattleShip struct {
	Points Set[Point]
}

func (b *BattleShip) AddPoint(p Point) {
	if b.Points == nil {
		b.Points = NewSet[Point]()
	}
	b.Points.Add(p)
}

func (b *BattleShip) RemovePoint(p Point) {
	b.Points.Remove(p)
}

func (b *BattleShip) Size() int {
	return b.Points.Size()
}

func (b *BattleShip) IsOneLine() bool {
	if len(b.Points) == 0 {
		return false
	}

	isOneLineX := true
	isOneLineY := true

	var prevPoint *Point
	for currentPoint, _ := range b.Points {
		if prevPoint == nil {
			p := currentPoint // создаём копию, чтобы взять адрес
			prevPoint = &p
			continue
		}
		if isOneLineX && prevPoint.X != currentPoint.X {
			isOneLineX = false
		}

		if isOneLineY && prevPoint.Y != currentPoint.Y {
			isOneLineY = false
		}

		if !isOneLineX && !isOneLineY {
			return false
		}

		p := currentPoint // создаём копию, чтобы взять адрес
		prevPoint = &p
		continue
	}

	return isOneLineX || isOneLineY //все координаты или на одной линии по X, или по Y
}

func (b *BattleShip) IsCorrect() bool {
	return 0 < b.Points.Size() && b.Points.Size() <= 4 && b.IsOneLine()
}
