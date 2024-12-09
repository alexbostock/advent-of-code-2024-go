package puzzle9

import (
	"fmt"
	"io"
)

type Input struct {
	Blocks []int
	Files  []File
}
type File struct {
	isEmpty bool
	id      int
	size    int
}

func ParseInput9(data io.Reader) Input {
	var blocks []int
	var files []File
	buffer := make([]byte, 100)

	currentByteIsFile := true
	currentFileID := 0
	for {
		n, err := data.Read(buffer)
		for _, digit := range buffer[:n] {
			num := int(digit) - '0'

			if currentByteIsFile {
				files = append(files, File{false, currentFileID, num})
				for i := 0; i < num; i++ {
					blocks = append(blocks, currentFileID)
				}
			} else {
				if num > 0 {
					files = append(files, File{true, 0, num})
				}
				for i := 0; i < num; i++ {
					blocks = append(blocks, -1)
				}
			}
			currentByteIsFile = !currentByteIsFile
			if !currentByteIsFile {
				currentFileID++
			}
		}
		if err == io.EOF {
			break
		}
	}
	return Input{blocks, files}
}

func parseByteAsDigit(char byte) (int, error) {
	num := int(char) - '0'
	if num < 0 || num > 9 {
		return 0, fmt.Errorf("unexpected character %v", char)
	}
	return num, nil
}

func MoveBlocks(blocks []int) {
	rightIndex := len(blocks) - 1
	for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
		if blocks[leftIndex] != -1 {
			continue
		}
		for rightIndex > leftIndex && blocks[rightIndex] == -1 {
			rightIndex--
		}
		if rightIndex > leftIndex {
			blocks[leftIndex] = blocks[rightIndex]
			blocks[rightIndex] = -1
		}
	}
}

func MoveFiles(files []File) []File {
	rightIndex := len(files) - 1
	for currentFileID := files[rightIndex].id; currentFileID >= 0; currentFileID-- {
		for rightIndex >= len(files) || files[rightIndex].isEmpty || files[rightIndex].id != currentFileID {
			rightIndex--
		}
		leftIndex, canMove := findSpaceForFile(files, rightIndex)
		if !canMove {
			continue
		}

		var filesNext []File
		for index, file := range files {
			if index == leftIndex {
				remainingSpace := files[leftIndex].size - files[rightIndex].size
				filesNext = append(filesNext, files[rightIndex])
				if remainingSpace > 0 {
					filesNext = append(filesNext, File{true, 0, remainingSpace})
				}
			} else if index == rightIndex {
				if filesNext[len(filesNext)-1].isEmpty {
					filesNext[len(filesNext)-1].size += files[rightIndex].size
				} else {
					filesNext = append(filesNext, File{true, 0, files[rightIndex].size})
				}
			} else if index == rightIndex+1 && file.isEmpty && files[rightIndex-1].isEmpty {
				filesNext[len(filesNext)-1].size += file.size
			} else {
				filesNext = append(filesNext, file)
			}
		}
		files = filesNext
	}
	return files
}

func findSpaceForFile(files []File, rightIndex int) (index int, canMoveFile bool) {
	for leftIndex := 0; leftIndex < rightIndex; leftIndex++ {
		if files[leftIndex].isEmpty && files[leftIndex].size >= files[rightIndex].size {
			return leftIndex, true
		}
	}
	return 0, false
}

func ComputeChecksum(blocks []int) int {
	sum := 0
	for index, fileID := range blocks {
		if fileID > -1 {
			sum += index * fileID
		}
	}
	return sum
}

func ComputeChecksumFiles(files []File) int {
	checksum := 0
	currentBlockIndex := 0
	for _, file := range files {

		for i := 0; i < file.size; i++ {
			if !file.isEmpty {
				checksum += currentBlockIndex * file.id
			}
			currentBlockIndex++
		}
	}
	return checksum
}
