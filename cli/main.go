package main

import (
	"encoding/json"
	"fmt"
	"github.com/jeyemwey/yuv4mpeg2"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}

	yuv, err := yuv4mpeg2.ParseHeader(file)
	if err != nil {
		log.Println(err)
	}

	j, _ := json.Marshal(yuv)
	fmt.Println(string(j))
}
