package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func write(w io.Writer, b []byte) {
	_, err := w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 512)
		for i := 1; i <= 5; i++ {
			chunk := fmt.Sprintf("This is chunk %d", i)
			write(w, []byte(chunk))

			rand.Read(buf)
			write(w, buf)
			w.(http.Flusher).Flush()

			time.Sleep(3 * time.Second)
		}
	})
	log.Fatal(http.ListenAndServe(":12345", nil))
}
