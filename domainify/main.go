package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

const allowedChars = "abcdefghijklmeopqrstuvwxyz0123456789_-"

func main() {
	var tlds []string
	flag.Parse()
	if flag.NArg() == 0 {
		tlds = []string{"com", "net"}
	} else {
		tlds = append(tlds, flag.Args()...)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
