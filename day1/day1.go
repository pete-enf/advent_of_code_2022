package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
	lines , err := os.Open("elf_snacks.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(lines)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	lines.Close()
	var lastVar int = 0
	for _, line := range text {
	    intVar, err := strconv.Atoi(line)
	    if intVar > lastVar {
		lastVar = intVar
	    }
	    fmt.Println(intVar, lastVar, err)
	}
}
