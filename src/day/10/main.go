package day10

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/10/input.txt")
	}
}

type TileType int16

const (
	Uknown TileType = iota
	Ground
	Start
	VerticalPipe
	HorizontalPipe
	NorthEastPipe
	NorthWestPipe
	SouthEastPipe
	SouthWestPipe
)

func (t TileType) String() string {
	switch t {
	case Ground:
		return "Ground"
	case Start:
		return "Start"
	case VerticalPipe:
		return "VerticalPipe"
	case HorizontalPipe:
		return "HorizontalPipe"
	case NorthEastPipe:
		return "NorthEastPipe"
	case NorthWestPipe:
		return "NorthWestPipe"
	case SouthEastPipe:
		return "SouthEastPipe"
	case SouthWestPipe:
		return "SouthWestPipe"
	default:
		return "invalid"
	}
}

func (t tile) String() string {
	return fmt.Sprintf("%v (%v, %v)", t.tileType, t.x, t.y)
}

var (
	tileTypeMap = map[string]TileType{
		".": Ground,
		"S": Start,
		"|": VerticalPipe,
		"-": HorizontalPipe,
		"L": NorthEastPipe,
		"J": NorthWestPipe,
		"7": SouthWestPipe,
		"F": SouthEastPipe,
	}
)

var (
	pipeDirections = map[TileType][]point{
		VerticalPipe:   []point{point{0, 1}, point{0, -1}},
		HorizontalPipe: []point{point{1, 0}, point{-1, 0}},
		NorthEastPipe:  []point{point{0, -1}, point{1, 0}},
		NorthWestPipe:  []point{point{0, -1}, point{-1, 0}},
		SouthEastPipe:  []point{point{0, 1}, point{1, 0}},
		SouthWestPipe:  []point{point{0, 1}, point{-1, 0}},
	}
)

type point struct {
	x int
	y int
}

type tile struct {
	id       string
	x        int
	y        int
	tileType TileType
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	tilesMatrix := parseMap(input)
	startTile := findStart(tilesMatrix)

	loopTiles := findLoopTiles(startTile, tilesMatrix)
	longestSteps := len(loopTiles) / 2
	return longestSteps, nil
}

func RunB(inputFile string) (int, error) {
	input := getInput(inputFile)
	tilesMatrix := parseMap(input)
	startTile := findStart(tilesMatrix)

	loopTiles := findLoopTiles(startTile, tilesMatrix)

	count := getEnclosedTilesCount(loopTiles, tilesMatrix)
	return count, nil
}

func parseMap(input []string) [][]tile {
	tilesMatrix := [][]tile{}
	for y, line := range input {
		tiles := []tile{}

		for x, tileString := range strings.Split(line, "") {
			tileType, ok := tileTypeMap[tileString]
			if !ok {
				log.Fatalf("invalid tile type: %v", tileString)
			}
			tiles = append(tiles, tile{
				id:       strconv.Itoa(x) + "x" + strconv.Itoa(y),
				x:        x,
				y:        y,
				tileType: tileType,
			})
		}

		tilesMatrix = append(tilesMatrix, tiles)
	}
	return tilesMatrix
}

func findStart(tiles [][]tile) tile {
	for _, line := range tiles {
		for _, tile := range line {
			if tile.tileType == Start {
				return tile
			}
		}
	}
	panic("no start found")
}

func findConnectingPipes(start tile, matrix [][]tile) []tile {
	connectingTiles := []tile{}
	connectingTiles = append(connectingTiles, safeGetTile(start.x, start.y-1, matrix))
	connectingTiles = append(connectingTiles, safeGetTile(start.x, start.y+1, matrix))
	connectingTiles = append(connectingTiles, safeGetTile(start.x-1, start.y, matrix))
	connectingTiles = append(connectingTiles, safeGetTile(start.x+1, start.y, matrix))
	return funk.Filter(connectingTiles, func(t tile) bool {
		return t != tile{} && t.tileType != Ground && pipeConnectsToPoint(t, point{start.x, start.y})
	}).([]tile)
}

func safeGetTile(x int, y int, matrix [][]tile) tile {
	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[y]) {
		return tile{}
	}
	return matrix[y][x]
}

func getNextConnectingPipe(prev tile, current tile, matrix [][]tile) tile {
	nextTile := getNextTile(prev, current, matrix)

	if (nextTile == tile{} || nextTile.tileType == Ground || !pipeConnectsToPoint(current, point{nextTile.x, nextTile.y})) {
		return tile{}
	}

	return nextTile
}

func getNextTile(previous tile, current tile, matrix [][]tile) tile {
	directions := pipeDirections[current.tileType]
	dir1 := directions[0]
	dir2 := directions[1]

	con1 := safeGetTile(current.x+dir1.x, current.y+dir1.y, matrix)
	con2 := safeGetTile(current.x+dir2.x, current.y+dir2.y, matrix)

	if previous.id == con1.id {
		return con2
	}

	if previous.id != con2.id {
		if previous.tileType == Start {
			return tile{}
		}
		fmt.Printf("current: %v\n", current)
		fmt.Printf("con1: %v\ncon2: %v\n", con1, con2)
		fmt.Printf("previous: %v\n", previous)
		log.Fatalf("Previous tile does not connect to current tile?")
	}

	return con1
}

func pipeConnectsToPoint(pipe tile, p point) bool {
	directions := pipeDirections[pipe.tileType]
	dir1 := directions[0]
	dir2 := directions[1]

	return (p.x == pipe.x+dir1.x && p.y == pipe.y+dir1.y) ||
		(p.x == pipe.x+dir2.x && p.y == pipe.y+dir2.y)
}

func findLoopTiles(start tile, tilesMatrix [][]tile) []tile {
	connectingPipes := findConnectingPipes(start, tilesMatrix)
	paths := [][]tile{}
	parsedPaths := []string{start.id}

	for _, connectedTile := range connectingPipes {
		paths = append(paths, []tile{start, connectedTile})
		parsedPaths = append(parsedPaths, connectedTile.id)
	}

	connectingTile := tile{}
	for (connectingTile == tile{}) {
		for i := 0; i < len(paths); i++ {
			path := paths[i]
			lenPath := len(path)

			currentTile := path[lenPath-1]
			previousTile := path[lenPath-2]

			if (currentTile == tile{}) {
				continue
			}

			nextTile := getNextConnectingPipe(previousTile, currentTile, tilesMatrix)

			if (nextTile == tile{}) {
				paths = append(paths[:i], paths[i+1:]...)
				i--
				continue
			}

			paths[i] = append(paths[i], nextTile)

			if funk.Contains(parsedPaths, nextTile.id) {
				fmt.Printf("found looping tile: %v\n", nextTile)
				connectingTile = nextTile
				break
			}
			parsedPaths = append(parsedPaths, nextTile.id)
		}
	}

	loopPaths := funk.Filter(paths, func(path []tile) bool {
		return funk.Last(path).(tile).id == connectingTile.id
	}).([][]tile)

	if len(loopPaths) != 2 {
		panic("len loopPaths should be 2")
	}

	lenFirstHalf := len(loopPaths[0])

	firstLoopTile := loopPaths[0][1]
	lastLoopTile := loopPaths[1][1]

	startDir1 := point{firstLoopTile.x - start.x, firstLoopTile.y - start.y}
	startDir2 := point{lastLoopTile.x - start.x, lastLoopTile.y - start.y}
	startingPipeType := findPipeType(startDir1, startDir2)

	loopPaths[0][0].tileType = startingPipeType

	return append(loopPaths[0], loopPaths[1][1:lenFirstHalf-1]...)
}

func findPipeType(dir1 point, dir2 point) TileType {
	for k, v := range pipeDirections {
		if funk.Contains(v, dir1) && funk.Contains(v, dir2) {
			return k
		}
	}

	panic("no pipe type found")
}

func getEnclosedTilesCount(loop []tile, matrix [][]tile) int {
	loopMatrix := make([][]tile, len(matrix))

	for _, tile := range loop {
		loopMatrix[tile.y] = append(loopMatrix[tile.y], tile)
	}

	total := 0
	for _, row := range loopMatrix {
		sort.SliceStable(row, func(i, j int) bool {
			return row[i].x < row[j].x
		})
		total += countEnclosedTiles(row)
	}
	return total
}

func countEnclosedTiles(row []tile) int {
	if len(row) == 0 {
		return 0
	}

	pipes := funk.Filter(row, func(t tile) bool {
		return t.tileType != HorizontalPipe
	}).([]tile)

	count := 0
	leni := len(pipes)

	current := tile{}
	prev := tile{}
	insideStart := -1

	for i := 0; i < leni; i++ {
		current = pipes[i]
		if i > 0 {
			prev = pipes[i-1]
		}

		switch current.tileType {
		case VerticalPipe:
			if insideStart == -1 {
				insideStart = i
				continue
			}
			insideStart = -1

			break
		case NorthEastPipe:
			if insideStart == -1 {
				insideStart = i
				continue
			}
			break
		case SouthEastPipe:
			if insideStart == -1 {
				insideStart = i
				continue
			}
			break
		case NorthWestPipe:
			if insideStart != -1 {
				if (prev.tileType == NorthEastPipe && insideStart == i-1) ||
					(prev.tileType == SouthEastPipe && insideStart < i-1) {
					insideStart = -1
					continue
				}
			}
			continue
		case SouthWestPipe:
			if insideStart != -1 {
				if (prev.tileType == SouthEastPipe && insideStart == i-1) ||
					(prev.tileType == NorthEastPipe && insideStart < i-1) {
					insideStart = -1
					continue
				}
			}
			continue
		default:
			log.Fatalf("invalid tile type: %v", current.tileType)
		}

		if (prev == tile{}) {
			continue
		}

		count += current.x - prev.x - 1
	}

	return count
}
