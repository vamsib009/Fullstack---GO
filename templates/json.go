package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

type Jsonresponce struct {
	Symbol struct {
		Symbol               string  `json:"symbol"`
		DxSymbol             string  `json:"dxSymbol"`
		Exchange             string  `json:"exchange"`
		IsoExchange          string  `json:"isoExchange"`
		BzExchange           string  `json:"bzExchange"`
		Type                 string  `json:"type"`
		Name                 string  `json:"name"`
		Description          string  `json:"description"`
		Sector               string  `json:"sector"`
		Industry             string  `json:"industry"`
		Open                 float64 `json:"open"`
		High                 float64 `json:"high"`
		Low                  float64 `json:"low"`
		BidPrice             float64 `json:"bidPrice"`
		AskPrice             float64 `json:"askPrice"`
		AskSize              int     `json:"askSize"`
		BidSize              int     `json:"bidSize"`
		Size                 int     `json:"size"`
		BidTime              int64   `json:"bidTime"`
		AskTime              int64   `json:"askTime"`
		LastTradePrice       float64 `json:"lastTradePrice"`
		LastTradeTime        int64   `json:"lastTradeTime"`
		Volume               int     `json:"volume"`
		Change               float64 `json:"change"`
		ChangePercent        float64 `json:"changePercent"`
		PreviousClosePrice   float64 `json:"previousClosePrice"`
		FiftyDayAveragePrice float64 `json:"fiftyDayAveragePrice"`
		FiftyTwoWeekHigh     float64 `json:"fiftyTwoWeekHigh"`
		FiftyTwoWeekLow      float64 `json:"fiftyTwoWeekLow"`
		MarketCap            int64   `json:"marketCap"`
		SharesOutstanding    int64   `json:"sharesOutstanding"`
		Pe                   float64 `json:"pe"`
		ForwardPE            float64 `json:"forwardPE"`
		DividendYield        float64 `json:"dividendYield"`
		PayoutRatio          float64 `json:"payoutRatio"`
	} `json:"GE"`
}

func init() {
	tpl = template.Must(template.ParseFiles("json.html"))
}

func main() {

	fmt.Println("sucess")
	http.HandleFunc("/", datahandler)
	http.ListenAndServe(":8070", nil)
	data := []Jsonresponce{}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}

func datahandler(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("http://careers-data.benzinga.com/rest/richquoteDelayed?symbols=GE")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var s Jsonresponce

	err = json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}

}
