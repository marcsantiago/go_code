package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"time"
)

var printf = fmt.Printf
var print = fmt.Println

func main() {

	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[:1] {
		go fetch(url, ch) //start goroutine
	}
	for range os.Args[:1] {
		print(<-ch) //recieve from channel ch
	}
	printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send the error to the channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak memory
	if err != nil {
		ch <- fmt.Sprint(err) // send the error to the channel
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
