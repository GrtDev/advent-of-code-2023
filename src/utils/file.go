package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filePath string) string {
	inputPath, err := filepath.Abs(filePath)
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func ReadLines(filePath string) []string {
	return ToLines(ReadFile(filePath))
}

func ReadParagraphs(filePath string) []string {
	return ToParagraphs(ReadFile(filePath))
}

func ToLines(content string) []string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.Trim(content, "\n")
	return strings.Split(content, "\n")
}

func ToParagraphs(content string) []string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.Trim(content, "\n")
	return strings.Split(content, "\n\n")
}

func ToRows(content string) []string {
	return ToLines(content)
}

func ToColumns(content string) []string {
	rows := ToRows(content)
	rowsLength := len(rows[0])
	columns := make([]string, rowsLength)

	for _, row := range rows {
		for x, value := range row {
			columns[x] += string(value)
		}
	}

	return columns
}
