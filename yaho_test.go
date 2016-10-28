package yaho

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {

	teststocks, err := GetStocks("YHOO, AAPL, MSFT")

	if err != nil {
		log.Println("error")
	} else {
		log.Println(teststocks.Query.Results.Quote[0].Symbol)
		log.Println(teststocks.Query.Results.Quote[0].LastTradePriceOnly)
		log.Println(teststocks.Query.Results.Quote[1].Symbol)
		log.Println(teststocks.Query.Results.Quote[1].LastTradePriceOnly)
		log.Println(teststocks.Query.Results.Quote[2].Symbol)
		log.Println(teststocks.Query.Results.Quote[2].LastTradePriceOnly)
	}

	teststock, err := GetSingleStocks("YHOO")

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(teststock.Query.Results.Quote.Name)
		log.Println(teststock.Query.Results.Quote.LastTradePriceOnly)
	}

}
