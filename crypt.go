package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type statusBlockChain struct {
	USD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"USD"`
	AUD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"AUD"`
	BRL struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"BRL"`
	CAD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CAD"`
	CHF struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CHF"`
	CLP struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CLP"`
	CNY struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CNY"`
	DKK struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"DKK"`
	EUR struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"EUR"`
	GBP struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"GBP"`
	HKD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HKD"`
	INR struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"INR"`
	ISK struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"ISK"`
	JPY struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"JPY"`
	KRW struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"KRW"`
	NZD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"NZD"`
	PLN struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"PLN"`
	RUB struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"RUB"`
	SEK struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SEK"`
	SGD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SGD"`
	THB struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"THB"`
	TWD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"TWD"`
}

type statusCryptoCompare struct {
	BTC int     `json:"BTC"`
	EUR float64 `json:"EUR"`
}

type statusCoinMarketCap []struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
	PriceEur         string `json:"price_eur"`
	Two4HVolumeEur   string `json:"24h_volume_eur"`
	MarketCapEur     string `json:"market_cap_eur"`
}

type statusCryptoNator struct {
	Ticker struct {
		Base    string `json:"base"`
		Target  string `json:"target"`
		Price   string `json:"price"`
		Volume  string `json:"volume"`
		Change  string `json:"change"`
		Markets []struct {
			Market string  `json:"market"`
			Price  string  `json:"price"`
			Volume float64 `json:"volume"`
		} `json:"markets"`
	} `json:"ticker"`
	Timestamp int    `json:"timestamp"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
}

//Timing duration
type Duration int64

var (
	//Set Timing variables
	latestReload  = time.Now()
	previousValue float64

	statusBC  statusBlockChain
	statusCC  statusCryptoCompare
	statusCMC statusCoinMarketCap
	statusCN  statusCryptoNator

	//Color values
	printMagenta   = color.New(color.FgMagenta)
	printRed       = color.New(color.FgRed)
	printWhite     = color.New(color.FgWhite)
	printHighBlue  = color.New(color.FgHiBlue)
	printHighGreen = color.New(color.FgHiGreen)
)

//Set Constants (Variables that never change)
const (

	//Time
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute

	//Api urls
	urlBC  = "https://blockchain.info/ticker"
	urlCC  = "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=BTC,EUR"
	urlCMC = "https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=EUR"
	urlCN  = "https://api.cryptonator.com/api/full/btc-eur"

	//Iteration time
	refresTime = time.Second * 15
)

/*--------------
Database Loading
--------------*/
/*
*Function used to load in the database
 */
func loadDatabase() {
	printRed.Println("Establishing Database Connection")
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

/*--------------
Crypto Loading
--------------*/

/*
*Function used to load in all the data on launch
 */
func loadCrypto() {

	//Set time ticker
	t := time.NewTicker(refresTime)

	//Start looping
	for {
		loadStatusAllCrypto()
		displayFunction()
		//When t.C receives, the loop will continue
		<-t.C
	}
}

/*
*Load all coin stats
 */
func loadStatusAllCrypto() {
	json.Unmarshal(getBytesByUrl(urlBC), &statusBC)
	json.Unmarshal(getBytesByUrl(urlCC), &statusCC)
	json.Unmarshal(getBytesByUrl(urlCMC), &statusCMC)
	json.Unmarshal(getBytesByUrl(urlCN), &statusCN)
}

/*
*Gets the bytes for a json url
 */
func getBytesByUrl(jsonUrl string) []byte {

	//Get respons form url
	resp, err := http.Get(jsonUrl)
	if err != nil {
		fmt.Println(err)
	}

	//Turn json into bytes
	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

/*--------------
TERMINAL DISPLAY
--------------*/
/*
*Function used to display in the terminal
 */
func displayFunction() {
	fmt.Println("\n")
	printMagenta.Println("\nGetting new Crypto Status")
	printWhite.Printf("Current time:%v", time.Now())
	displayStatusBlockchain()
	displayStatusCryptocompare()
	displayStatusCoinmarketcap()
	displayStatusCryptonator()
	fmt.Println("\n")
}

func displayStatusBlockchain() {
	printHighBlue.Println("\nBlockchain status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusBC.EUR.Recent)
}

func displayStatusCryptocompare() {
	printHighBlue.Println("\nCryptocompare status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCC.EUR)
}

func displayStatusCoinmarketcap() {
	printHighBlue.Println("\nCoinmarketcap status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCMC[0].PriceEur)
}

func displayStatusCryptonator() {
	printHighBlue.Println("\nCryptonator status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCN.Ticker.Price)
}

/*--------------
WEB DISPLAY
--------------*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", statusBC.EUR.Recent)
}

/*
*MAIN FUNCTION
 */
func main() {
	go loadDatabase()
	go loadCrypto()
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
