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
var newestComic comicInfo

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
				newestComic = newComic
				comicInfoMap[newestComic.Number] = newestComic
				for i := 1; i < 121; i++ {
					url := "http://xkcd.com/" + strconv.Itoa(newComic.Number-i) + "/info.0.json"
					go fetch(url, ch)
				}
				for i := 1; i < 121; i++ {
					newComic := <-ch
					comicInfoMap[newComic.Number] = newComic
				}
				for i := 121; i < 241; i++ {
					url := "http://xkcd.com/" + strconv.Itoa(newComic.Number-i) + "/info.0.json"
					go fetch(url, ch)
				}
				for i := 121; i < 241; i++ {
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

	var encoder = json.NewEncoder(w)
	if r.URL.Path == "/" {
		encoder.Encode(newestComic)
	} else {
		res := strings.Split(r.URL.Path, "/")
		path, err := strconv.Atoi(res[1])
		if err != nil {
			return
		}
		if path < 11 {
			number := path
			fmt.Println(number)
			newComicInfoSlice := []comicInfo{}
			for i := 0; i < 10; i++ {
				index := newestComic.Number - i - 10*number
				newComicInfoSlice = append(newComicInfoSlice, comicInfoMap[index])
			}
			encoder.Encode(newComicInfoSlice)
		}
	}
}
