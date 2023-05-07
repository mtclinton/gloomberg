package main

import "fmt"
import "net/http"
import "log"
import "io/ioutil"
type Client struct {
    client http.Client
    base string
}

func newClient() Client {
    return Client {
        client: http.Client{},
        base: "https://query1.finance.yahoo.com",
    }
}

func main() {
    fmt.Println("Gloomberg")
    c:= newClient()
    req, err := http.NewRequest(http.MethodGet, c.base+"/v10/finance/quoteSummary/aapl?modules=earningsHistory", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "gloomberg")

	res, getErr := c.client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
    if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
    fmt.Println(body)
}
