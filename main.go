package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%d : %q\n", lineNumber, line)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR")
		return
	}

	//testtingif file open works before adding other parts
	fmt.Println("OK")
}
