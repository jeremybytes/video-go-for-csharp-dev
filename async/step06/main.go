package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}
	}
	defer resp.Body.Close()
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	return p
}

type person struct {
	ID           int
	GivenName    string
	FamilyName   string
	StartDate    time.Time
	Rating       int
	FormatString string
}

func (p person) String() string {
	if p.FormatString != "" {
		return fmt.Sprintf("%s %s", p.FamilyName, p.GivenName)
	}
	return fmt.Sprintf("%s %s", p.GivenName, p.FamilyName)
}

func main() {
	start := time.Now()
	ids := getIds()
	fmt.Println(ids)

	if len(os.Args) > 1 && os.Args[1] == "noasync" {
		for _, n := range ids {
			p := getPerson(n)
			fmt.Printf("%d: %v\n", p.ID, p)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %s\n", elapsed)
}
