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
		for i := 0; i < len(teststocks.Query.Results.Quote); i++ {
			log.Println("Hämtar flera stocks [")
			log.Println(teststocks.Query.Results.Quote[i].Symbol)
			log.Println(teststocks.Query.Results.Quote[i].LastTradePriceOnly)
		}
	}

	teststock, err := GetSingleStocks("YHOO")

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Hämtar en stocks [")
		log.Println(teststock.Query.Results.Quote.Name)
		log.Println(teststock.Query.Results.Quote.LastTradePriceOnly)
	}

}
