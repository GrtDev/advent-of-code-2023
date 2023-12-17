package day14

import (
	"advent-of-code-2023/old/matrix"
	mat "advent-of-code-2023/old/matrix"
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadParagraphs(inputFile)
	} else {
		return utils.ReadParagraphs("./day/14/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	platformMap := matrix.Create2D(input[0])

	rocks := getRocksNorthIndex(platformMap)
	total := countNorthLoad(rocks, platformMap)
	return total, nil
}

func RunB(inputFile string) (int, error) {
	input := getInput(inputFile)
	platformMap := matrix.Create2D(input[0])

	rockIndexes, cubeIndexes := getIndexMatrices(platformMap)

	rocks := mat.IndexMatrix{
		Indexes: rockIndexes,
		Width:   mat.NumColumns(platformMap),
		Height:  mat.NumRows(platformMap),
	}
	cubes := mat.IndexMatrix{
		Indexes: cubeIndexes,
		Width:   mat.NumColumns(platformMap),
		Height:  mat.NumRows(platformMap),
	}

	for i := 0; i < 1000; i++ {
		rocks, cubes = cycleRocks(rocks, cubes)
	}

	total := countNorthLoad(rocks.Indexes, platformMap)
	return total, nil
}

func countNorthLoad(rocks [][]int, platformMap [][]string) int {
	total := 0
	lenColumn := matrix.LenColumn(platformMap)
	for _, col := range rocks {
		for _, rockI := range col {
			total += lenColumn - rockI
		}
	}
	return total
}

func getIndexMatrices(m [][]string) ([][]int, [][]int) {
	rockRegexp := regexp.MustCompile(`O`)
	cubeRegexp := regexp.MustCompile(`#`)
	columns := matrix.ColumnStrings(m)

	rockIndexes := make([][]int, matrix.NumColumns(m))
	cubeIndexes := make([][]int, matrix.NumColumns(m))

	for x, column := range columns {
		rockIndexes[x] = utils.TakeFirstInts(rockRegexp.FindAllStringIndex(column, -1))
		cubeIndexes[x] = utils.TakeFirstInts(cubeRegexp.FindAllStringIndex(column, -1))
	}
	return rockIndexes, cubeIndexes
}

func getRocksNorthIndex(m [][]string) [][]int {
	northIndexes := make([][]int, matrix.NumColumns(m))
	columns := matrix.ColumnStrings(m)
	rockIndexes, cubeIndexes := getIndexMatrices(m)
	for x, _ := range columns {
		northIndexes[x] = moveRocksLeft(rockIndexes[x], cubeIndexes[x])
	}
	return northIndexes
}

func memoizeKey(rocks mat.IndexMatrix, cubes mat.IndexMatrix) string {
	r := strings.Join(utils.IntsToString(mat.ValuesIndexes(rocks)), ",")
	c := strings.Join(utils.IntsToString(mat.ValuesIndexes(cubes)), ",")
	return fmt.Sprintf("%s-%v", r, c)
}

var countVariantsMap = map[string]mat.IndexMatrix{}

func cycleRocks(rocks mat.IndexMatrix, cubes mat.IndexMatrix) (mat.IndexMatrix, mat.IndexMatrix) {
	key := memoizeKey(rocks, cubes)
	if collectVariantCount, ok := countVariantsMap[key]; ok {
		return collectVariantCount, cubes
	}

	// north
	rocks.Indexes = moveAllRocksLeft(rocks.Indexes, cubes.Indexes)

	// west
	rocks = mat.RotateIndexRight(rocks)
	cubes = mat.RotateIndexRight(cubes)
	rocks.Indexes = moveAllRocksLeft(rocks.Indexes, cubes.Indexes)

	// south
	rocks = mat.RotateIndexRight(rocks)
	cubes = mat.RotateIndexRight(cubes)
	rocks.Indexes = moveAllRocksLeft(rocks.Indexes, cubes.Indexes)

	// east
	rocks = mat.RotateIndexRight(rocks)
	cubes = mat.RotateIndexRight(cubes)
	rocks.Indexes = moveAllRocksLeft(rocks.Indexes, cubes.Indexes)

	// set back to north
	rocks = mat.RotateIndexRight(rocks)
	cubes = mat.RotateIndexRight(cubes)
	return rocks, cubes
}

func moveAllRocksLeft(rocks [][]int, cubes [][]int) [][]int {
	for x := 0; x < len(rocks); x++ {
		rockCol := rocks[x]
		cubesCol := cubes[x]
		rocks[x] = moveRocksLeft(rockCol, cubesCol)
	}
	return rocks
}

func moveRocksLeft(rocks []int, cubes []int) []int {
	movedRockIndexes := []int{}
	newI, lastI := -1, -1
	lenCubes := len(cubes)

	for _, rockI := range rocks {
		if lenCubes == 0 || rockI < cubes[0] {
			if lastI == -1 {
				newI = 0
			} else {
				newI = lastI + 1
			}
		} else {
			for i := 0; i < lenCubes; i++ {
				currCubeI := cubes[i]
				nextCubeI := -1

				if i < len(cubes)-1 {
					nextCubeI = cubes[i+1]
				}

				if rockI > currCubeI && rockI < nextCubeI || nextCubeI == -1 {
					newI = currCubeI + 1
					break
				}

				if i == lenCubes-2 {
					newI = nextCubeI + 1
					break
				}
			}
		}

		if newI <= lastI {
			newI = lastI + 1
		}
		movedRockIndexes = append(movedRockIndexes, newI)
		lastI = newI
	}

	return movedRockIndexes
}

func printMap(rockColumnIndexes [][]int, platformMap [][]string) {
	platformMap = matrix.ReplaceString("O", ".", platformMap)
	for x, col := range rockColumnIndexes {
		for _, y := range col {
			if platformMap[x][y] != "." {
				platformMap[x][y] = "X"
			} else {
				platformMap[x][y] = "O"
			}
		}
	}
	log.Printf("%+v\n", matrix.Print(platformMap))
}

func printMapB(rocks [][]int, cubes [][]int, platformMap [][]string) {
	platformMap = matrix.ReplaceString("O", ".", platformMap)

	for x, col := range cubes {
		for _, y := range col {
			if platformMap[x][y] != "#" {
				platformMap[x][y] = "X"
			} else {
				platformMap[x][y] = "#"
			}
		}
	}

	for x, col := range rocks {
		for _, y := range col {
			if platformMap[x][y] != "." {
				platformMap[x][y] = "X"
			} else {
				platformMap[x][y] = "O"
			}
		}
	}

	log.Printf("%+v\n", matrix.Print(platformMap))
}
