package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for i, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		filename := fmt.Sprintf("file%d", i)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer out.Close()
		if _, err := io.Copy(out, resp.Body); err != nil {
			log.Fatal(err)
		}
	}
}

// cd 到当前目录下
// go run main.go <your-url>