# xkcdReader

Online demo: [site](http://davidguan.me/xkcdReader/)    
Mobile application built with React Native: [Github repo](https://github.com/EcutDavid/react-native-xkcdReader)    

![image](https://cloud.githubusercontent.com/assets/10692276/16587760/bce8a7c0-42fe-11e6-9154-b67909e30284.png)


### Why you need a Golang data API.
* Becasue of CORS, can't get the data provided by xkcd directly in front-end.
* So we can package the data to 10 comics per request. Instead of 1 request match one comic.

### How to develop based on this repo
Clone this repo   
`go run Api/main.go`   
`cd Client/ && npm i && npm start`   
Now do whatever you want.

Data from https://xkcd.com/json.html.
