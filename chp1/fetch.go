package main

import (
	_ "bufio"
	"fmt"
	"io"
	_ "io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	print := fmt.Println
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			print("Error")
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		b := os.Stdout
		n, err := io.Copy(b, resp.Body)
		status := resp.Status
		resp.Body.Close()
		if err != nil {
			print("Error")
		}
		fmt.Printf("%s\n", n)
		print(status)
	}
}
