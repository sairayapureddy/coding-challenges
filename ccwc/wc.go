package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	bytePtr := flag.Bool("c", false, "count the number of bytes")
	linePtr := flag.Bool("l", false, "count the number of lines")
	wordPtr := flag.Bool("w", false, "count the number of lines")
	charPtr := flag.Bool("m", false, "count the number of characters")
	flag.Parse()

	fileName := ""
	args := flag.Args()
	if len(args) == 1 {
		fileName = flag.Args()[0]
	}

	var data string
	if fileName != "" {
		bytes, ok := os.ReadFile(fileName)
		if ok == nil {
			data = string(bytes)
		} else {
			fmt.Println(ok.Error())
		}
	} else {
		bytes := make([]byte, 0)
		reader := bufio.NewReader(os.Stdin)
		reader.Read(bytes)
		data = string(bytes)
	}

	if *bytePtr {
		fmt.Printf("%d %s \n", len(data), fileName)
	}
	if *linePtr {
		fmt.Printf("%d %s \n", len(strings.Split(data, "\n")), fileName)
	}
	if *wordPtr {
		pattern := regexp.MustCompile(`\s+`)
		fmt.Printf("%d %s \n", len(pattern.FindAllString(data, -1)), fileName)
	}
	if *charPtr {
		fmt.Printf("%d %s \n", utf8.RuneCountInString(data), fileName)
	}
}