package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getIds() (ids []int) {
	resp, err := http.Get("http://localhost:9874/people/ids")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(body, &ids)
	return
}

func main() {
	ids := getIds()
	fmt.Println(ids)
}
