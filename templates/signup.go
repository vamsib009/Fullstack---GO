package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error
var tpl *template.Template

type JsonResponse map[string]struct {
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
}

var username, password string

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username = req.FormValue("username")
	password = req.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		//res.Write([]byte("User created!"))
		http.ServeFile(res, req, "search.html")
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/", 301)
		return
	}
	http.ServeFile(res, req, "search.html")
	fmt.Println(databaseUsername)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "home.html")
}

func landpage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "search.html")
	er := tpl.ExecuteTemplate(res, "search.html", username)
	if er != nil {
		panic(er)
	}
}
func init() {
	tpl = template.Must(template.ParseFiles("json.html", "search.html"))
}
func datahandler(w http.ResponseWriter, r *http.Request) {

	x := r.URL.Query()
	y := x["symbol"]
	res, error := http.Get("http://careers-data.benzinga.com/rest/richquoteDelayed?symbols=" + y[0])
	if error != nil {
		panic(error)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var s JsonResponse

	err = json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}

	er := tpl.ExecuteTemplate(w, "json.html", s)
	if er != nil {
		panic(er)
	}
}

func main() {
	db, err = sql.Open("mysql", "stock:Vamsi1994@tcp(stock.citsm1ymtilj.us-east-1.rds.amazonaws.com:3306)/stock")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected database")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/json", datahandler)
	http.HandleFunc("/land", landpage)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)

}
