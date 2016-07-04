# xkcdReader

Fetch the data provided by https://xkcd.com/json.html, then provide the data to front end via json.
![image](https://cloud.githubusercontent.com/assets/10692276/16563967/96d828ca-4237-11e6-9911-1b7d6af8df70.png)


### Why you need a Golang data API.
* Becasue of CORS, can't get the data provided by xkcd directly in front-end.
* So we can package the data to 10 comics per request. Instead of 1 request match one comic.

### How to develop based on this repo
Clone this repo   
`go run Api/main.go`   
`cd Client/ && npm i && npm start`
