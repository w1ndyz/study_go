package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"study_go/section4/4.5/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	github.SortResult(result)

	var month, year, distant []*github.Issue
	for v, i := range result.Items {
		if i.CreatedAt.After(time.Now().Add(time.Hour * 24 * 30 * -1)) {
			month = append(month, result.Items[v])
		} else if i.CreatedAt.After(time.Now().Add(time.Hour * 24 * 365 * -1)) {
			year = append(year, result.Items[v])
		} else {
			distant = append(distant, result.Items[v])
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("with in a month\n")
	for _, item := range month {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("with in a year\n")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Printf("long long ago\n")
	for _, item := range distant {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
