package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getIds() (ids []int, err error) {
	resp, err := http.Get("http://localhost:9874/people/ids")
	if err != nil {
		return nil, fmt.Errorf("error fetching IDs: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading IDs: %v", err)
	}
	err = json.Unmarshal(body, &ids)
	if err != nil {
		return nil, fmt.Errorf("error parsing IDs: %v", err)
	}
	return ids, nil
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

func getPersonChannel(id int, ch chan<- person) {
	ch <- getPerson(id)
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
	ids, err := getIds()
	if err != nil {
		// fmt.Printf("getIDs failed: %v", err)
		// os.Exit(1)
		log.Fatalf("getIDs failed: %v", err)
	}
	fmt.Println(ids)

	if len(os.Args) > 1 && os.Args[1] == "noasync" {
		for _, n := range ids {
			p := getPerson(n)
			fmt.Printf("%d: %v\n", p.ID, p)
		}
	} else {
		ch := make(chan person)
		for _, n := range ids {
			go getPersonChannel(n, ch)
		}
		for range ids {
			p := <-ch
			fmt.Printf("%d: %v\n", p.ID, p)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %s\n", elapsed)
}
