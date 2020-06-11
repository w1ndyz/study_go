package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type database struct {
	R     map[string]dollars
	mutex sync.Mutex
}

type dollars float32

func main() {
	db := database{map[string]dollars{"shoes": 50, "socks": 5}, sync.Mutex{}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db.R {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprint(w, "invalid price")
	} else {
		db.R[item] = dollars(p)
		fmt.Fprint(w, "update price success")
	}
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := r.URL.Query().Get("item")
	delete(db.R, item)
	fmt.Fprint(w, "delete success")
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if price, ok := db.R[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
