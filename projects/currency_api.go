package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// function call
	// API_KEY = 1a4ddeafc2e385d5ccb13920573329cc (FREE SUBSCRIPTION)
	getByCurrency("YOUR-API-KEY", "EUR", "TRY")
	getSpecificDateCurrency("YOUR-API-KEY", "2013-12-24", "EUR", "TRY,USD")
	getDateRangeCurrency("YOUR-API-KEY", "2012-05-01", "2012-05-25", "EUR", "TRY,USD")
}

// getByCurrency allow you to make a api call for based on currency and expected result
func getByCurrency(apiKey, base, symbol string) {
	var apiCall = "http://api.exchangeratesapi.io/v1/latest?access_key=" + apiKey + "&base=" + base + "&symbols=" + symbol
	response, errorResponse := http.Get(apiCall)
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}
	//We Read the response body on the line below.
	apiCall = apiCall + "&" + base + "&" + symbol
	body, errorResponse := ioutil.ReadAll(response.Body)
	//Convert the body to type string
	callResponse := string(body)
	log.Printf(callResponse)
}

// getSpecificDateCurrency allow you to make a api call for based on currency and expected result and allow you to pass specific date
func getSpecificDateCurrency(apiKey, date, base, symbol string) {
	var apiCall = "http://api.exchangeratesapi.io/v1/" + date + "?access_key=" + apiKey + "&base=" + base + "&symbols=" + symbol
	response, errorResponse := http.Get(apiCall)
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}
	//We Read the response body on the line below.
	apiCall = apiCall + "&" + base + "&" + symbol
	body, errorResponse := ioutil.ReadAll(response.Body)
	//Convert the body to type string
	callResponse := string(body)
	log.Printf(callResponse)
}

// getDateRangeCurrency allow you to make a api call for based on currency and expected result and allow you to pass date range
func getDateRangeCurrency(apiKey, startDate, endDate, base, symbol string) {
	var apiCall = "http://api.exchangeratesapi.io/v1/timeseries?access_key=" + apiKey + "&base=" + base + "&symbols=" + symbol + "&start_date" + startDate + "&end_date" + endDate
	response, errorResponse := http.Get(apiCall)
	if errorResponse != nil {
		log.Fatalln(errorResponse)
	}
	//We Read the response body on the line below.
	apiCall = apiCall + "&" + base + "&" + symbol
	body, errorResponse := ioutil.ReadAll(response.Body)
	//Convert the body to type string
	callResponse := string(body)
	log.Printf(callResponse)
}
