package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

func main() {
	file, err := os.Open(`/Users/t-ninomiya/go/src/github.com/miliya12/cli-oreilly/sprinkle/src.txt`)
	if err != nil {
		fmt.Errorf("ERROR: %s", err.Error)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	rand.Seed(time.Now().UTC().UnixNano())
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		t := lines[rand.Intn(len(lines))]
		fmt.Println(strings.Replace(t, otherWord, sc.Text(), -1))
	}
}
