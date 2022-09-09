package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	_, _ = fmt.Fprintf(w, "<body style=\"background-color: #111;\">")
	_, _ = fmt.Fprintf(w, "<h1 style=\""+
		"color:#eee; font-family:'Verdana',sans-serif; font-size: 275px; font-weight: bold; "+
		"letter-spacing: -1px; line-height: 1; text-align: center;  "+
		"\">Void</h1>")

	_, _ = fmt.Fprintf(w, "<h2 style=\""+
		"color:#eee; font-family:'Helvetica Neue',sans-serif; font-size: 30px; font-weight: 300; "+
		"line-height: 32px; margin: 0 0 72px; text-align: center;  "+
		"\">There is nothing here... really nothing, just void</h2>")
	_, _ = fmt.Fprintf(w, "</body>")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	_, _ = fmt.Fprintf(w, "Hello!")
}
