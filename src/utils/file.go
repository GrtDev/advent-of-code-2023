package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile( filePath string)string {
    inputPath, err := filepath.Abs(filePath)
    input, err := os.ReadFile(inputPath)
    if(err != nil) { log.Fatal(err) }
    return string(input)
}

func ReadLines(filePath string)[]string {
    return ToLines(ReadFile(filePath))
}

func ToLines(content string)[]string {
    content = strings.ReplaceAll(content, "\r\n", "\n")
    content = strings.Trim(content, "\n")
    return strings.Split(content, "\n")
}


