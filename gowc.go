package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	cFlg := flag.Bool("c", false, "print the byte counts")
	lFlg := flag.Bool("l", false, "print the character counts")
	wFlg := flag.Bool("w", false, "print the newline counts")
	mFlg := flag.Bool("m", false, "print the word counts")
	hFlg := flag.Bool("h", false, "display help information")

	flag.Usage = func() {
		fmt.Println(`Usage: ccwc [FLAG] [FILE PATH]

Flag Usage:

  -c    print the byte counts
  -l    print the character counts
  -m    print the word counts
  -w    print the newline counts
  -h    display help information`)
	}
	flag.Parse()

	if *hFlg {
		flag.Usage()
		os.Exit(0)
	}

	var file *os.File
	if len(flag.Args()) == 0 {
		file = os.Stdin
	} else {
		f, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file = f
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	fileData := []byte{}
	for scanner.Scan() {
		fileData = append(fileData, scanner.Bytes()...)
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
		os.Exit(1)
	}

	if *cFlg {
		fmt.Println(len(fileData), file.Name())
	} else if *lFlg {
		fmt.Println(bytes.Count(fileData, []byte("\n")), file.Name())
	} else if *wFlg {
		fmt.Println(len(bytes.Fields(fileData)), file.Name())
	} else if *mFlg {
		fmt.Println(len(bytes.Split(fileData, []byte(""))), file.Name())
	} else {
		fmt.Println(bytes.Count(
			fileData, []byte("\n")),
			len(bytes.Fields(fileData)),
			len(fileData),
			file.Name(),
		)
	}

	os.Exit(0)
}
