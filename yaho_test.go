package yaho

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	log.Println("Testing Get")
	teststring := "YHOO"
	var tst Stock
	tst, err := Get(teststring)

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(tst.Query.Results.Quote.Symbol)
	}

}

func TestPolyGet(t *testing.T) {
	log.Println("Testing PolyGet")
	teststrings := "YHOO, APPL, MSFT"
	var tst Stocks
	tst, err := PolyGet(teststrings)

	if err != nil {
		log.Println("error")
	} else {
		log.Println(tst.Query.Results.Quote[0].Symbol)
		log.Println(tst.Query.Results.Quote[1].Symbol)
		log.Println(tst.Query.Results.Quote[2].Symbol)
	}

}
