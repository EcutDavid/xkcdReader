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

func init() {
	ch := make(chan comicInfo)
	if res, err := http.Get("http://xkcd.com/info.0.json"); err == nil {
		if res.StatusCode == http.StatusOK {
			newComic := comicInfo{}
			if err := json.NewDecoder(res.Body).Decode(&newComic); err == nil {
				newestComic = newComic
				comicInfoMap[newestComic.Number] = newestComic
				for j := 0; j < 9; j++ {
					addtion := 100 * j
					for i := addtion + 1; i < addtion+101; i++ {
						indexStr := strconv.Itoa(newComic.Number - i)
						url := "http://xkcd.com/" + indexStr + "/info.0.json"
						go fetch(url, ch)
					}
					for i := addtion + 1; i < addtion+101; i++ {
						newComic := <-ch
						comicInfoMap[newComic.Number] = newComic
					}
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9021", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept")

	var encoder = json.NewEncoder(w)
	if strings.Index(r.URL.Path, "/search") != -1 {
		res := strings.Split(r.URL.Path, "/")
		r.Body.Close()
		if len(res) == 3 && res[2] != "" {
			newComicInfoSlice := []comicInfo{}
			counter := 0
			for _, value := range comicInfoMap {
				if strings.Index(strings.ToLower(value.Title), strings.ToLower(res[2])) != -1 {
					newComicInfoSlice = append(newComicInfoSlice, value)
					counter++
					if counter > 5 {
						break
					}
				}
			}
			encoder.Encode(newComicInfoSlice)
		}
		return
	}

	if r.URL.Path == "/" {
		encoder.Encode(newestComic)
	} else {
		res := strings.Split(r.URL.Path, "/")
		r.Body.Close()
		path, err := strconv.Atoi(res[1])
		if err != nil {
			return
		}
		if path < 100 {
			number := path
			newComicInfoSlice := []comicInfo{}
			for i := 0; i < 10; i++ {
				index := newestComic.Number - i - 10*number
				newComicInfoSlice = append(newComicInfoSlice, comicInfoMap[index])
			}
			encoder.Encode(newComicInfoSlice)
		}
	}
}
