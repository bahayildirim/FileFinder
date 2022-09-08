package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var wordToCheck string
	var pathToCheck string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input path to check: ")
	scanner.Scan()
	pathToCheck = scanner.Text()
	fmt.Println("Input word to check: ")
	fmt.Scan(&wordToCheck)
	matchCount := 0

	dirFunc := func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, wordToCheck) {
			matchCount++
		}
		return nil
	}

	filepath.WalkDir(pathToCheck, dirFunc)

	checkedWord := flag.String("word", wordToCheck, "Word to find in directories")
	foundWords := flag.Int("matches", matchCount, "Displays the amount of matches with the given word")

	flag.Parse()

	fmt.Println("Checked word: ", *checkedWord)
	fmt.Println("Amount of matches: ", *foundWords)
}

/*
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := make(chan int)
	go typeNum(c)

	for i := 0; i < 1000; i++ {
		c <- i
	}
	close(c)

	wg.Wait()
}

func typeNum(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
	wg.Done()
}
*/
