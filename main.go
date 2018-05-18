package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	seperator string
	findStr   string
	path      string
)

func init() {
	flag.StringVar(&seperator, "sp", "", "block seperator default line feed '\\n'")
	flag.StringVar(&path, "path", "", "find path default .")
	flag.StringVar(&findStr, "find", "", "find string")

	flag.Parse()

	seperator = strings.Replace(seperator, "\\n", "\n", -1)
	seperator = strings.Replace(seperator, "\\r", "\r", -1)
	seperator = strings.Replace(seperator, "\\t", "\t", -1)

	fmt.Println("find:", findStr)
	fmt.Println("path:", path)
	fmt.Println("sp:", seperator, len(seperator))
}

func main() {
	if findStr == "" {
		panic("invalid find parametar")
	}
	if seperator == "" {
		seperator = "\n"
	}
	if path == "" {
		path = "."
	}

	if err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			fmt.Println("===== Search.....", path, "=====")

			scanner := bufio.NewScanner(file)

			scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				if atEOF && len(data) == 0 {
					return 0, nil, nil
				}
				if i := strings.Index(string(data), seperator); i >= 0 {
					return i + len(seperator), data[0:i], nil
				}
				if atEOF {
					return len(data), data, nil
				}
				// Request more data.
				return 0, nil, nil
			})

			for scanner.Scan() {
				if strings.Contains(scanner.Text(), findStr) {
					fmt.Println(scanner.Text())
				}
			}
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}
