package main

import "os"

func main() {
	worklist := make(chan []string)  //可能有重复的url列表
	unseenLinks := make(chan string) //去重后的url列表

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundlinks := crawl(link)
				go func() {
					worklist <- foundlinks
				}()
			}
		}()
	}

	// 主goroutine对URL列表进行去重
	// 并把没有爬取过的条目发送给爬虫程序
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
