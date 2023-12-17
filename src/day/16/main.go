package day16

import (
	"advent-of-code-2023/utils"
	"advent-of-code-2023/utils/matrix"
	"fmt"
	"slices"
	"strings"

	"github.com/thoas/go-funk"
)

func getInput(inputFile string) string {
	if inputFile != "" {
		return utils.ReadFile(inputFile)
	} else {
		return utils.ReadFile("./day/16/input.txt")
	}
}

type point struct {
	x, y int
}

type movement struct {
	x, y      int
	direction *point
}

type tile struct {
	value         string
	hitCount      int
	hitDirections []string
}

func (mov *movement) move() {
	mov.x += mov.direction.x
	mov.y += mov.direction.y
}

func (mov *movement) String() string {
	return fmt.Sprintf("x: %v, y: %v, direction: %+v", mov.x, mov.y, mov.direction)
}

func (p *point) String() string {
	return fmt.Sprintf("(%v,%v)", p.x, p.y)
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	matrix := createMatrix(input)

	total := calcTilesHit([]*movement{
		{0, 0, &point{1, 0}},
	}, matrix)

	return total, nil
}

func RunB(inputFile string) (int, error) {

	input := getInput(inputFile)
	matrix := createMatrix(input)

	startingPoints := []*movement{}

	for x := 0; x < matrix.Width; x++ {
		startingPoints = append(startingPoints,
			&movement{
				x:         x,
				y:         0,
				direction: &point{x: 0, y: 1},
			},
			&movement{
				x:         x,
				y:         matrix.Width - 1,
				direction: &point{x: 0, y: -1},
			},
		)

	}

	for y := 0; y < matrix.Height; y++ {
		startingPoints = append(startingPoints,
			&movement{
				x:         0,
				y:         y,
				direction: &point{x: 1, y: 0},
			},
			&movement{
				x:         matrix.Width - 1,
				y:         y,
				direction: &point{x: -1, y: 0},
			},
		)
	}

	totals := []int{}

	for _, start := range startingPoints {
		totals = append(totals, calcTilesHit([]*movement{start}, matrix))
	}

	max := funk.MaxInt(totals)

	return max, nil
}

func createMatrix(input string) matrix.Matrix[*tile] {
	lines := utils.ToLines(input)
	rows := make([][]*tile, len(lines))

	for y, line := range lines {
		values := strings.Split(line, "")
		rows[y] = make([]*tile, len(values))
		for x, v := range values {
			rows[y][x] = &tile{
				value:         v,
				hitCount:      0,
				hitDirections: []string{},
			}
		}
	}

	return matrix.Create(rows)
}

func resetMatrix(matrix matrix.Matrix[*tile]) {
	for _, tile := range matrix.AllValues() {
		tile.hitCount = 0
		tile.hitDirections = []string{}
	}
}

func calcTilesHit(movements []*movement, matrix matrix.Matrix[*tile]) int {

	resetMatrix(matrix)

	for len(movements) > 0 {
		for i := 0; i < len(movements); i++ {
			move := movements[i]
			dir := move.direction
			dirId := fmt.Sprint(*dir)
			tile := matrix.Get(move.x, move.y)

			if tile == nil || slices.Contains(tile.hitDirections, dirId) {
				movements = append(movements[:i], movements[i+1:]...)
				i--
				continue
			}

			tile.hitCount++
			tile.hitDirections = append(tile.hitDirections, dirId)

			switch tile.value {
			case "|":
				if move.direction.x != 0 {
					move.direction.x = 0
					move.direction.y = -1
					movements = append(movements, &movement{
						x:         move.x,
						y:         move.y,
						direction: &point{x: 0, y: 1},
					})

				}
				break
			case "-":
				if move.direction.y != 0 {
					move.direction.x = -1
					move.direction.y = 0
					movements = append(movements, &movement{
						x:         move.x,
						y:         move.y,
						direction: &point{x: 1, y: 0},
					})

				}
				break
			case "\\":
				move.direction.x = move.direction.x ^ move.direction.y
				move.direction.y = move.direction.x ^ move.direction.y
				move.direction.x = move.direction.x ^ move.direction.y
				break
			case "/":
				move.direction.x = move.direction.x ^ move.direction.y
				move.direction.y = move.direction.x ^ move.direction.y
				move.direction.x = move.direction.x ^ move.direction.y
				move.direction.x *= -1
				move.direction.y *= -1
				break
			}

			move.move()
		}
	}

	total := 0
	for _, tile := range matrix.AllValues() {
		if tile.hitCount > 0 {
			total++
		}
	}

	return total
}
