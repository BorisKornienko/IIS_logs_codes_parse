package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	forParse, err := ioutil.ReadDir("ForParse")
	if err != nil {
		log.Fatal(err)
	}

	type FileCode struct {
		Code  string
		Value int
	}
	type AllFilesCode struct {
		File  string
		Codes []FileCode
	}

	var allOutput AllFilesCode

	for _, file := range forParse {
		codeMap := make(map[string]int)
		f, err := os.Open("ForParse/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			statusCode := strings.Split(scanner.Text(), " ")
			if len(statusCode) > 4 {
				code := statusCode[len(statusCode)-4]
				// fmt.Println(code)
				codeMap[code]++
			}
		}

		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		for key, value := range codeMap {
			var fileCode FileCode
			fileCode.Code = key
			fileCode.Value = value
			allOutput.File = file.Name()
			// println(allOutput.File)
			allOutput.Codes = append(allOutput.Codes, fileCode)
		}
	}

	var jsonData []byte
	jsonData, err = json.Marshal(allOutput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
