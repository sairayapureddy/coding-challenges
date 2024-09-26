package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
		totalBytes := make([]byte, 0)
		reader := bufio.NewReader(os.Stdin)
		for {
			bytes, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			totalBytes = append(totalBytes, bytes...)
		}
		data = string(totalBytes)
	}

	if *bytePtr {
		fmt.Printf("%d %s \n", len(data), fileName)
	}
	if *linePtr {
		fmt.Printf("%d %s \n", len(strings.Split(data, "\n"))-1, fileName)
	}
	if *wordPtr {
		pattern := regexp.MustCompile(`\s+`)
		fmt.Printf("%d %s \n", len(pattern.FindAllString(data, -1)), fileName)
	}
	if *charPtr {
		fmt.Printf("%d %s \n", utf8.RuneCountInString(data), fileName)
	}

	if !*bytePtr && !*linePtr && !*wordPtr && !*charPtr {
		pattern := regexp.MustCompile(`\s+`)
		fmt.Printf("%d %d %d %s \n", len(strings.Split(data, "\n"))-1, len(pattern.FindAllString(data, -1)), len(data), fileName)
	}
}