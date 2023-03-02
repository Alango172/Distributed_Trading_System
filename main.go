package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		panic("The input is not valid, please insert the right file path")
	}

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	filescanner := bufio.NewScanner(readFile)
	filescanner.Split(bufio.ScanLines)
	var fileLines []string

	for filescanner.Scan() {
		fileLines = append(fileLines, filescanner.Text())
	}

	defer readFile.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}

}
