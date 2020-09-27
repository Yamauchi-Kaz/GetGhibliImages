package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var ghibliurl string = "http://www.ghibli.jp/gallery/"
	// var movieTitle = "marnie"		//思い出のマーニー（2014）
	// var movieTitle = "kaguyahime"	//かぐや姫の物語（2013）
	// var movieTitle = "kazetachinu" 	//風立ちぬ（2013）
	// var movieTitle = "kokurikozaka" 	//コクリコ坂から（2011）
	// var movieTitle = "karigurashi" 		//借りぐらしのアリエッティ（2010）
	// var movieTitle = "ponyo" 			//崖の上のポニョ（2008）
	// var movieTitle = "ged" 				//ゲド戦記（2006）
	var movieTitle = "chihiro" //千と千尋の神隠し（2001）

	var filename = ""
	var imageURL = ""

	for i := 1; i <= 50; i++ {
		if i < 10 {
			filename = movieTitle + "00" + strconv.Itoa(i) + ".jpg"
		} else if i >= 10 && i <= 50 {
			filename = movieTitle + "0" + strconv.Itoa(i) + ".jpg"
		} else {
			break
		}

		imageURL = ghibliurl + filename
		response, err := http.Get(imageURL)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		io.Copy(file, response.Body)
	}

}
