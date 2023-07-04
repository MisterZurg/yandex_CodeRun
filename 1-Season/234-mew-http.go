package main

import (
	"fmt"
	"net/http"
)

const URL = "http://127.0.0.1:7777/"

func getDataFromHttpService(in []string) string {
	goWaySet := make(map[string]bool)
	fourHeaders := []string{}
	for _, v := range in {
		if _, ok := goWaySet[v]; ok {
			continue
		}
		goWaySet[v] = true
		fourHeaders = append(fourHeaders, v)
	}

	req, _ := http.NewRequest("MEW", URL, nil)
	req.Header = http.Header{
		"X-Cat-Variable": fourHeaders,
	}
	client := &http.Client{}
	resp, _ := client.Do(req)
	data := resp.Header.Get("X-Cat-Value")
	return data
}

func getValues(data string) []string {

}

func mewHTTP(in []string) string {
	data := getDataFromHttpService(in)
	values := getValues(data)
	//
	//// X-Cat-Value
	//
	//fmt.Println(resp.Header.Get("X-Cat-Value"))
}

func main() {
	requiredVariables := make([]string, 4)
	for i := range requiredVariables {
		fmt.Scan(&requiredVariables[i])
	}

	mewHTTP(requiredVariables)

	//req, err := http.NewRequest("GET", URL, nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//req.Header = http.Header{
	//	"X-Cat-Variable": {a, b, c, d},
	//}
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//
	//// X-Cat-Value
	//
	//fmt.Println(resp.Header.Get("X-Cat-Value"))
}
