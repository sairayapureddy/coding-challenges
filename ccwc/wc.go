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

type commandLineOptions struct {
	countBytes, countLines, countWords, countChars bool
}

func main() {
	var options commandLineOptions
	flag.BoolVar(&options.countBytes, "c", false, "count the number of bytes")
	flag.BoolVar(&options.countLines, "l", false, "count the number of lines")
	flag.BoolVar(&options.countWords, "w", false, "count the number of lines")
	flag.BoolVar(&options.countChars, "m", false, "count the number of characters")
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

	if options.countBytes {
		fmt.Printf("%d %s \n", len(data), fileName)
	}
	if options.countLines {
		fmt.Printf("%d %s \n", len(strings.Split(data, "\n"))-1, fileName)
	}
	if options.countWords {
		pattern := regexp.MustCompile(`\s+`)
		fmt.Printf("%d %s \n", len(pattern.FindAllString(data, -1)), fileName)
	}
	if options.countChars {
		fmt.Printf("%d %s \n", utf8.RuneCountInString(data), fileName)
	}

	if !options.countBytes &&
		!options.countLines &&
		!options.countWords &&
		!options.countChars {
		pattern := regexp.MustCompile(`\s+`)
		fmt.Printf("%d %d %d %s \n", len(strings.Split(data, "\n"))-1, len(pattern.FindAllString(data, -1)), len(data), fileName)
	}
}