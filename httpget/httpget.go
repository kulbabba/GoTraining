package main

import ("encoding/json" 
	"net/http"
	"fmt")

func main() {
	// TODO: write content of the page to stdout
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		
	}


	result := make(map[string]string)
	dec := json.NewDecoder(resp.Body)
	error := dec.Decode(&result)
	if error != nil{
		
	}
	i := result["origin"]
	fmt.Println("My ip address is ", i)

}
