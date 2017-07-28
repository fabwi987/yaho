# yaho

A package for retreiving financial data to your application.
The package uses the [Yahoo Finance API](https://finance.yahoo.com/) to fetch the latest financial data.

https://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quote%20where%20symbol%20in%20(%22YHOO%22%2C%22AAPL%22%2C%22GOOG%22%2C%22MSFT%22)&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys

https://developer.yahoo.com/yql/console/?q=show%20tables&env=store://datatables.org/alltableswithkeys#h=select+*+from+yahoo.finance.quote+where+symbol+in+(%22YHOO%22%2C%22AAPL%22%2C%22GOOG%22%2C%22MSFT%22)

##Example

Use the following command for a single or multiple quote:
```Go
yagoo.Get("YHOO")
yagoo.PolyGet("YHOO, MSFT, ...")
```

Types
```Go
type Stock struct {
	Query struct {
		Count   int       `json:"count"`
		Created time.Time `json:"created"`
		Lang    string    `json:"lang"`
		Results struct {
			Quote Quote `json:"quote"`
		} `json:"results"`
	} `json:"query"`
}

type Stocks struct {
	Query struct {
		Count   int       `json:"count"`
		Created time.Time `json:"created"`
		Lang    string    `json:"lang"`
		Results struct {
			Quote []Quote `json:"quote"`
		} `json:"results"`
	} `json:"query"`
}

type Quote struct {
	Symbol                                         string      `json:"symbol"`
	Ask                                            string      `json:"Ask"`
	AverageDailyVolume                             string      `json:"AverageDailyVolume"`
	Bid                                            string      `json:"Bid"`
	AskRealtime                                    interface{} `json:"AskRealtime"`
	BidRealtime                                    interface{} `json:"BidRealtime"`
	BookValue                                      string      `json:"BookValue"`
	ChangePercentChange                            string      `json:"Change_PercentChange"`
	Change                                         string      `json:"Change"`
	Commission                                     interface{} `json:"Commission"`
	Currency                                       string      `json:"Currency"`
	ChangeRealtime                                 interface{} `json:"ChangeRealtime"`
	AfterHoursChangeRealtime                       interface{} `json:"AfterHoursChangeRealtime"`
	DividendShare                                  string      `json:"DividendShare"`
	LastTradeDate                                  string      `json:"LastTradeDate"`
	TradeDate                                      interface{} `json:"TradeDate"`
	EarningsShare                                  string      `json:"EarningsShare"`
	ErrorIndicationreturnedforsymbolchangedinvalid interface{} `json:"ErrorIndicationreturnedforsymbolchangedinvalid"`
	EPSEstimateCurrentYear                         string      `json:"EPSEstimateCurrentYear"`
	EPSEstimateNextYear                            string      `json:"EPSEstimateNextYear"`
	EPSEstimateNextQuarter                         string      `json:"EPSEstimateNextQuarter"`
	DaysLow                                        string      `json:"DaysLow"`
	DaysHigh                                       string      `json:"DaysHigh"`
	YearLow                                        string      `json:"YearLow"`
	YearHigh                                       string      `json:"YearHigh"`
	HoldingsGainPercent                            interface{} `json:"HoldingsGainPercent"`
	AnnualizedGain                                 interface{} `json:"AnnualizedGain"`
	HoldingsGain                                   interface{} `json:"HoldingsGain"`
	HoldingsGainPercentRealtime                    interface{} `json:"HoldingsGainPercentRealtime"`
	HoldingsGainRealtime                           interface{} `json:"HoldingsGainRealtime"`
	MoreInfo                                       interface{} `json:"MoreInfo"`
	OrderBookRealtime                              interface{} `json:"OrderBookRealtime"`
	MarketCapitalization                           string      `json:"MarketCapitalization"`
	MarketCapRealtime                              interface{} `json:"MarketCapRealtime"`
	EBITDA                                         string      `json:"EBITDA"`
	ChangeFromYearLow                              string      `json:"ChangeFromYearLow"`
	PercentChangeFromYearLow                       string      `json:"PercentChangeFromYearLow"`
	LastTradeRealtimeWithTime                      interface{} `json:"LastTradeRealtimeWithTime"`
	ChangePercentRealtime                          interface{} `json:"ChangePercentRealtime"`
	ChangeFromYearHigh                             string      `json:"ChangeFromYearHigh"`
	PercebtChangeFromYearHigh                      string      `json:"PercebtChangeFromYearHigh"`
	LastTradeWithTime                              string      `json:"LastTradeWithTime"`
	LastTradePriceOnly                             string      `json:"LastTradePriceOnly"`
	HighLimit                                      interface{} `json:"HighLimit"`
	LowLimit                                       interface{} `json:"LowLimit"`
	DaysRange                                      string      `json:"DaysRange"`
	DaysRangeRealtime                              interface{} `json:"DaysRangeRealtime"`
	FiftydayMovingAverage                          string      `json:"FiftydayMovingAverage"`
	TwoHundreddayMovingAverage                     string      `json:"TwoHundreddayMovingAverage"`
	ChangeFromTwoHundreddayMovingAverage           string      `json:"ChangeFromTwoHundreddayMovingAverage"`
	PercentChangeFromTwoHundreddayMovingAverage    string      `json:"PercentChangeFromTwoHundreddayMovingAverage"`
	ChangeFromFiftydayMovingAverage                string      `json:"ChangeFromFiftydayMovingAverage"`
	PercentChangeFromFiftydayMovingAverage         string      `json:"PercentChangeFromFiftydayMovingAverage"`
	Name                                           string      `json:"Name"`
	Notes                                          interface{} `json:"Notes"`
	Open                                           string      `json:"Open"`
	PreviousClose                                  string      `json:"PreviousClose"`
	PricePaid                                      interface{} `json:"PricePaid"`
	ChangeinPercent                                string      `json:"ChangeinPercent"`
	PriceSales                                     string      `json:"PriceSales"`
	PriceBook                                      string      `json:"PriceBook"`
	ExDividendDate                                 string      `json:"ExDividendDate"`
	PERatio                                        string      `json:"PERatio"`
	DividendPayDate                                string      `json:"DividendPayDate"`
	PERatioRealtime                                interface{} `json:"PERatioRealtime"`
	PEGRatio                                       string      `json:"PEGRatio"`
	PriceEPSEstimateCurrentYear                    string      `json:"PriceEPSEstimateCurrentYear"`
	PriceEPSEstimateNextYear                       string      `json:"PriceEPSEstimateNextYear"`
	Symbol1                                        string      `json:"Symbol1"`
	SharesOwned                                    interface{} `json:"SharesOwned"`
	ShortRatio                                     string      `json:"ShortRatio"`
	LastTradeTime                                  string      `json:"LastTradeTime"`
	TickerTrend                                    interface{} `json:"TickerTrend"`
	OneyrTargetPrice                               string      `json:"OneyrTargetPrice"`
	Volume                                         string      `json:"Volume"`
	HoldingsValue                                  interface{} `json:"HoldingsValue"`
	HoldingsValueRealtime                          interface{} `json:"HoldingsValueRealtime"`
	YearRange                                      string      `json:"YearRange"`
	DaysValueChange                                interface{} `json:"DaysValueChange"`
	DaysValueChangeRealtime                        interface{} `json:"DaysValueChangeRealtime"`
	StockExchange                                  string      `json:"StockExchange"`
	DividendYield                                  string      `json:"DividendYield"`
	PercentChange                                  string      `json:"PercentChange"`
}
```
