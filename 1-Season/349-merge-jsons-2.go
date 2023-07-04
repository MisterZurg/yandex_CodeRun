//package main
//
//import (
//	"bufio"
//	"encoding/json"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Offers struct {
//	Offers []Offer `json:"offers"`
//}
//
//type Offer struct {
//	OfferID    string `json:"offer_id"`
//	MarketSKU  int    `json:"market_sku"`
//	Price      int    `json:"price"`
//	ofIDnumber int    // helper
//}
//
//func main() {
//	sc := bufio.NewScanner(os.Stdin)
//	//
//	sc.Scan()
//	line := sc.Text()
//	nm := strings.Split(line, " ")
//	n, _ := strconv.Atoi(nm[0])
//	m, _ := strconv.Atoi(nm[1])
//	//susJson := []string{
//	//	`{"offers": [{"offer_id": "offer4", "market_sku": 10846332, "price": 1490}, {"offer_id": "offer2", "market_sku": 682644, "price": 499}]}`,
//	//	`{"offers": [{"offer_id": "offer3", "market_sku": 832784, "price": 14000}, {"offer_id": "offer1", "market_sku": 3234, "price": 100}]}`,
//	//}
//	////for i := range susJson {
//	////	fmt.Println(susJson[i])
//	////}
//	feeds := Offers{}
//	////fmt.Println(n, m)
//	//n := len(susJson)
//	for i := 0; i < n; i++ {
//		sc.Scan()
//		var currOffer Offers
//		susJson := sc.Bytes()
//		json.Unmarshal(susJson, &currOffer)
//		//json.Unmarshal([]byte(susJson[i]), &currOffer)
//		feeds.Offers = append(feeds.Offers, currOffer.Offers...)
//	}
//	//fmt.Println(feeds)
//	feeds.fillIds()
//	feeds.sortOffersById()
//	a, _ := json.Marshal(Offers{feeds.Offers[:m]})
//	fmt.Println(string(a))
//
//}
//
//func (ofs *Offers) fillIds() {
//	for i, of := range ofs.Offers {
//		idSuffix := strings.TrimPrefix(of.OfferID, "offer")
//		fmt.Println(idSuffix)
//		ofs.Offers[i].ofIDnumber, _ = strconv.Atoi(idSuffix)
//	}
//}
//func (ofs *Offers) sortOffersById() {
//	sort.Slice(ofs.Offers, func(i, j int) bool {
//		//iIdx := strings.TrimPrefix(ofs.Offers[i].OfferID, "offer")
//		//jIdx := strings.TrimPrefix(ofs.Offers[j].OfferID, "offer")
//		//
//		//return ofs.Offers[i].OfferID < ofs.Offers[j].OfferID
//		return ofs.Offers[i].ofIDnumber < ofs.Offers[j].ofIDnumber
//	})
//}
//
//// {"offers": [{"offer_id": "offer1", "market_sku": 10846332, "price": 1490}, {"offer_id": "offer2", "market_sku": 682644, "price": 499}]}
//// {"offers": [{"offer_id": "offer3", "market_sku": 832784, "price": 14000}, {"offer_id": "offer4", "market_sku": 3234, "price": 100}]}
//
////birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
////var bird Bird
////json.Unmarshal([]byte(birdJson), &bird)
////fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)

//package main
//
//// Made with https://www.codeconvert.ai/python-to-golang-converter
//import (
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//func main() {
//	var n, m int
//	fmt.Scan(&n, &m)
//	allRows := map[string]interface{}{
//		"offers": []interface{}{},
//	}
//	for i := 0; i < n; i++ {
//		var feed map[string]interface{}
//		json.NewDecoder(os.Stdin).Decode(&feed)
//		allRows["offers"] = append(allRows["offers"].([]interface{}), feed["offers"].([]interface{})...)
//	}
//	allRows["offers"] = allRows["offers"].([]interface{})[:m]
//	jsonData, _ := json.Marshal(allRows)
//	fmt.Println(string(jsonData))
//}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	feed := make(map[string]interface{})
	feed["offers"] = []interface{}{}
	offers := feed["offers"].([]interface{})
	var s string
	fmt.Scanln(&s)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanln(&s)
		var parsed map[string]interface{}
		json.Unmarshal([]byte(s), &parsed)
		for _, offer := range parsed["offers"].([]interface{}) {
			if len(offers) < m {
				offers = append(offers, offer)
			}
		}
	}
	jsonData, _ := json.Marshal(feed)
	fmt.Println(string(jsonData))
}

/*
// From pastebin C++
#include <iostream>
#include <algorithm>

#include "json.hpp"

int main() {
    std::ios_base::sync_with_stdio(false); std::cin.tie(0); std::cout.tie(0);
    int n, m;
    std::cin >> n >> m;
    nlohmann::json feed;
    feed["offers"] = nlohmann::json::array();
    auto& offers = feed["offers"];

    std::string s;
    std::getline(std::cin, s);
    for (int i = 0; i < n; i++) {
        std::string s;
        std::getline(std::cin, s);
        nlohmann::json parsed = nlohmann::json::parse(s);
        for (const auto& offer : parsed["offers"]) {
            if (offers.size() < m) {
                offers.push_back(offer);
            }
        }
    }

    std::cout << feed << std::endl;
    return 0;
}
*/
