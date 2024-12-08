package puzzle4

import (
	"bufio"
	"io"
)

type WordSearch struct {
	chars  [][]rune
	height int
	width  int
}

func ParseInput4(data io.Reader) WordSearch {
	var chars [][]rune
	scanner := bufio.NewScanner(data)
	height := 0
	width := 0
	for scanner.Scan() {
		chars = append(chars, []rune(scanner.Text()))
		height++
		width = len(scanner.Text())
	}
	return WordSearch{
		chars, height, width,
	}
}

func CountXMASInWordSearch(wordSearch WordSearch) int {
	count := 0
	for i := 0; i < wordSearch.height; i++ {
		for j := 0; j < wordSearch.width; j++ {
			patterns := [][]struct {
				i int
				j int
			}{
				{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}},
				{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}},
				{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}},
				{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}},
				{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}},
				{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}},
				{{i, j}, {i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}},
				{{i, j}, {i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}},
			}
			for _, pattern := range patterns {
				if hasXMASInWordSearchAtPositions(wordSearch, pattern) {
					count++
				}
			}
		}
	}
	return count
}

func hasXMASInWordSearchAtPositions(wordSearch WordSearch, coords []struct {
	i int
	j int
}) bool {
	for pos, coord := range coords {
		if coord.i < 0 || coord.i >= wordSearch.height || coord.j < 0 || coord.j >= wordSearch.width {
			return false
		}
		expected := rune("XMAS"[pos])
		if wordSearch.chars[coord.i][coord.j] != expected {
			return false
		}
	}
	return true
}

func CountCrossMASInWordSearch(wordSearch WordSearch) int {
	count := 0
	for i := 0; i < wordSearch.height; i++ {
		for j := 0; j < wordSearch.width; j++ {
			if hasCrossMASInWordSearchCentredAt(wordSearch, i, j) {
				count++
			}
		}
	}
	return count
}

func hasCrossMASInWordSearchCentredAt(wordSearch WordSearch, i, j int) bool {
	if i-1 < 0 || i+1 >= wordSearch.height || j-1 < 0 || j+1 >= wordSearch.width {
		return false
	}
	topLeft := wordSearch.chars[i-1][j-1]
	topRight := wordSearch.chars[i-1][j+1]
	bottomLeft := wordSearch.chars[i+1][j-1]
	bottomRight := wordSearch.chars[i+1][j+1]
	centre := wordSearch.chars[i][j]
	diagonalOneOkay := topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M'
	diagonalTwoOkay := topRight == 'M' && bottomLeft == 'S' || topRight == 'S' && bottomLeft == 'M'
	return centre == 'A' && diagonalOneOkay && diagonalTwoOkay
}
