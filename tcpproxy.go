package main

import (
	"io"
	"net/http"
	"log"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

//ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections. 
//Accepted connections are configured to enable TCP keep-alives. 
//Handler is typically nil, in which case the DefaultServeMux is used.