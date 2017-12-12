package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetJsonFile(str string) map[string]string {
	resp, err := http.Get(str)
	if err != nil {
		println("unable to load data from ", str, err.Error())
		return nil
	}

	result := make(map[string]string)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	if err != nil {
		println("unable to parse json", err.Error())
		return nil
	}
	return result
}
func GetMyIp(myMap map[string]string) {
	ip, ok := myMap["origin"]
	if !ok {
		fmt.Println("ip not found at map:", myMap)
		return
	}
	fmt.Println("My ip address is ", ip)
}
func main() {

	json := GetJsonFile("http://httpbin.org/uuid")
	GetMyIp(json)

}
