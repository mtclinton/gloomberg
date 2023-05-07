package main

type Chart struct {
	chart ChartStatus `json:"chart"`
}

type ChartStatus struct {
	result	[]ChartData `json:"result,omitempty"`
	chart_error Error `json:"error,omitempty"`
}

type Error struct {
	code        string `json:"code"`
	description string `json:"description"`
}

type ChartData struct {
	chartMeta ChartMeta `json:"meta"`
	timestamp  []int `json:"timestamp"`
	indicators Indicators `json:"indicators"`
}

type ChartMeta struct {
	currency             string  `json:"currency"`
	symbol               string  `json:"symbol"`
	exchangeName         string  `json:"exchangeName"`
	instrumentType       string  `json:"instrumentType"`
	firstTradeDate       int     `json:"firstTradeDate"`
	regularMarketTime    int     `json:"regularMarketTime"`
	gmtoffset            int     `json:"gmtoffset"`
	timezone             string  `json:"timezone"`
	exchangeTimezoneName string  `json:"exchangeTimezoneName"`
	regularMarketPrice   float64 `json:"regularMarketPrice"`
	chartPreviousClose   float64 `json:"chartPreviousClose"`
	previousClose        float64 `json:"previousClose"`
	scale                int     `json:"scale"`
	priceHint            int     `json:"priceHint"`
	currentTradingPeriod struct {
		Pre struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"pre"`
		Regular struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"regular"`
		Post struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"post"`
	} `json:"currentTradingPeriod"`
	tradingPeriods [][]struct {
		Timezone  string `json:"timezone"`
		Start     int    `json:"start"`
		End       int    `json:"end"`
		Gmtoffset int    `json:"gmtoffset"`
	} `json:"tradingPeriods"`
	dataGranularity string   `json:"dataGranularity"`
	crange           string   `json:"range"`
	validRanges     []string `json:"validRanges"`
}

type Indicators struct {
	quote []ChartQuote `json:"quote"`
}

type ChartQuote struct {
	qclose  []float64 `json:"close"`
	volume []int     `json:"volume"`
	low    []float64 `json:"low"`
	open   []float64 `json:"open"`
	high   []float64 `json:"high"`
}



