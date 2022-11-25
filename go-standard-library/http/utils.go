package main

const BASE_URL = "https://httpbin.org"

type Person struct {
	Name       string   `json:"name"`
	Address    string   `json:"addr"`
	Age        int      `json:"-"`
	FaveColors []string `json:"favs,omitempty"`
}
