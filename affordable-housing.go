package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
func main() {
	fmt.Println("starting the golang application")
	neighborhood :="Treasure Island"
	safeNeighborhood := url.QueryEscape(neighborhood)
	// affordable beds in Treasure Island
	url1 := fmt.Sprintf("https://data.sfgov.org/resource/yd5s-bd6e.json?$where=affordable_beds>0&$select=sum(affordable_beds)&neighborhood=%s",safeNeighborhood)
	fmt.Println("url1: ", url1)
	req, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(data))
	}

	url2 := fmt.Sprintf("https://data.sfgov.org/resource/yd5s-bd6e.json?$where=homeless_units>0&$select=sum(homeless_units)")
	fmt.Println("url2: ", url2)
	req2, err2 := http.NewRequest("GET", url2, nil)
	if err2 != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err2)
		return
	}

	//client := &http.Client{}
	resp2, err2 := client.Do(req2)

	if err2 != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err2)
		return
	} else {
		data, _ := ioutil.ReadAll(resp2.Body)
        fmt.Println(string(data))
	}
}