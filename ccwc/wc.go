package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	bytePtr := flag.Bool("c", false, "count the number of bytes")
	flag.Parse()

	fileName := ""
	args := flag.Args()
	if len(args) == 1 {
		fileName = flag.Args()[0]
	}

	var byteCount int
	if fileName != "" {
		data, ok := os.ReadFile(fileName)
		if ok == nil {
			byteCount = len(data)
		} else {
			fmt.Println(ok.Error())
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		byteCount = reader.Size()
	}

	if *bytePtr {
		fmt.Printf("%d %s \n", byteCount, fileName)
	}
}