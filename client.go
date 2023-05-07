package main

import "net/http"

type Version string

const (
    V7 Version = "v7"
    V8 Version = "v8"
    V10 Version = "v10"
)

type Client struct {
    client http.Client
    base string
}

func newClient() Client {
    return Client {
        client: http.Client{}
        base: "https://query1.finance.yahoo.com"
}


