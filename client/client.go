package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:12346", nil)
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Set("Accept-Encoding", "gzip")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("All response headers:")
	for k, v := range resp.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	var reader io.Reader
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("Error creating gzip reader:", err)
			return
		}
		defer gzipReader.Close()
		reader = gzipReader
	} else {
		reader = resp.Body
	}

	buf := make([]byte, 512)
	chunk := 0
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("chunk: %d, read %d bytes\n", chunk, n)
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("chunk: %d, read %d bytes\n", chunk, n)
		chunk++
	}
}
