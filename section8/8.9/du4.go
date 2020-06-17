package main

import (
	"os"
	"sync"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	// 当检测到输入是取消遍历
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for {
		select {
		case <-done:
			for range filesizes {

			}
			return
		case size, ok := <-filesizes:

		}
	}
}

func walkDir(dir string, n *sync.WaitGroup, filseSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		// ...
	}
}
