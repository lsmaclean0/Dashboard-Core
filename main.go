package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// stock struct
type stock struct {
	ID    string `json:"id"`
	Price int
}

func main() {

	// temporary, API key will be stored in config file
	apikey := "d12454c7cdmshefda0fe212b6c8bp12d401jsnbfd1685cf68c"

	// Generic tickers, will change to user input one frontend is ready
	tickers := []string{"SNAP", "TSLA"}

	for i := 0; i < len(tickers); i++ {
		ticker := tickers[i]
		log.Println("Getting prices for ticker: ", ticker)
		getTickers(ticker, apikey)
	}

	fmt.Println("Compiled succesfully")
}

func getTickers(ticker string, apikey string) {

	url := "https://realstonks.p.rapidapi.com/" + ticker

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", apikey)
	req.Header.Add("X-RapidAPI-Host", "realstonks.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

}
