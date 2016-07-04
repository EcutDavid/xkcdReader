package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type comicInfo struct {
	Number int    `json:"num"`
	Img    string `json:"img"`
	Title  string `json:"title"`
	Alt    string `json:"alt"`
}

var comicInfoMap = make(map[int]comicInfo)

func fetch(url string, ch chan<- comicInfo) {
	if res, err := http.Get(url); err == nil {
		if res.StatusCode == http.StatusOK {
			newComic := comicInfo{}
			if err := json.NewDecoder(res.Body).Decode(&newComic); err == nil {
				fmt.Println(newComic)
				ch <- newComic
			}
		}
	}
}

func main() {
	ch := make(chan comicInfo)
	if res, err := http.Get("http://xkcd.com/info.0.json"); err == nil {
		if res.StatusCode == http.StatusOK {
			newComic := comicInfo{}
			if err := json.NewDecoder(res.Body).Decode(&newComic); err == nil {
				comicInfoMap[newComic.Number] = newComic
				for i := 1; i < 51; i++ {
					url := "http://xkcd.com/" + strconv.Itoa(newComic.Number-i) + "/info.0.json"
					go fetch(url, ch)
				}
				for i := 1; i < 51; i++ {
					newComic := <-ch
					comicInfoMap[newComic.Number] = newComic
				}
				for i := 51; i < 101; i++ {
					url := "http://xkcd.com/" + strconv.Itoa(newComic.Number-i) + "/info.0.json"
					go fetch(url, ch)
				}
				for i := 51; i < 101; i++ {
					newComic := <-ch
					comicInfoMap[newComic.Number] = newComic
				}
			}
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9020", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept")

	method := r.Method
	var encoder = json.NewEncoder(w)
	if strings.Index(r.URL.Path, "/books") == 0 {
		if r.URL.Path == "/books" && method == "GET" {
			encoder.Encode(comicInfoMap)
		}
	}
}
