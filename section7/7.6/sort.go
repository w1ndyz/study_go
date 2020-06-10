package main

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type customSort struct {
	t []*Track
	r []string
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var c = customSort{tracks, []string{}}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	for i := len(x.r) - 1; i >= 0; i-- {
		switch x.r[i] {
		case "Title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "Artist":
			if x.t[i].Artist != x.t[j].Artist {
				return x.t[i].Artist < x.t[j].Artist
			}
		case "Album":
			if x.t[i].Artist != x.t[j].Artist {
				return x.t[i].Artist < x.t[j].Artist
			}
		case "Year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		default:
			panic("Not support field")
		}
	}
	return false
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func Sort(s string) {
	c.r = append(c.r, s)
	if len(c.r) > 4 {
		c.r = c.r[:4]
	}
	sort.Sort(c)
}

func main() {
	Sort("Title")
	for _, v := range tracks {
		fmt.Println(*v)
	}
	fmt.Println("=======")
	Sort("Artist")
	for _, v := range tracks {
		fmt.Println(*v)
	}
	fmt.Println("=======")
	Sort("Album")
	for _, v := range tracks {
		fmt.Println(*v)
	}
}
