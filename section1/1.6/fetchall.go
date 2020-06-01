package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https") {
			url = "https://" + url
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// nbytes, err := io.Copy(ioutil.WriteFile(), resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	ch <- fmt.Sprintf("while reading %s: %v", url, err)
	// 	return
	// }
	bytes, err := ioutil.ReadAll(resp.Body)
	data := []byte(bytes)
	fileName := "website_body_" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	if ioutil.WriteFile(fileName, data, 0644) == nil {
		fmt.Println("写入文件成功！文件名是:", fileName)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
