/*
There are two methods to decoding a JSON,

One is by loading the data into a declared struct
the other is by loading it into a map.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Story struct {
	Id         string   `json:"Id"`
	Title      string   `json:"Title"`
	Characters []string `json:"characters"`
}

func main() {
	data := []byte(`
		{
			"id": "12he34",
			"title": "Three dwarves",
			"characters": [
				"sneezy",
				"sleepy",
				"sleazy"
			]
		}
	`)

	// Unmarshal Json into a struct
	var s Story
	err := json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Title)
	fmt.Println(s.Characters)

	// Unmarshal Json into a map
	var m map[string]interface{}
	err2 := json.Unmarshal(data, &m)
	if err2 != nil {
		panic(err)
	}
	fmt.Println(m["title"])
	fmt.Println(m["characters"])
}
