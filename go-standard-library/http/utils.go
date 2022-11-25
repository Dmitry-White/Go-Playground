package main

const BASE_URL = "https://httpbin.org"

type Person struct {
	// XMLName    xml.Name `xml:"persondata"`
	Name       string   `json:"name" xml:"firstname"`
	Address    string   `json:"addr" xml:"address"`
	Age        int      `json:"-" xml:"age,attr"`
	FaveColors []string `json:"favs,omitempty" xml:"favecolors"`
}
