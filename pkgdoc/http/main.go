package main

import (
	"fmt"
	"net/http"
	"net/url"
)


func main()  {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})
	url.Parse()
	fmt.Println(http.ParseHTTPVersion("HTTP/1.0"))
	http.ListenAndServe(":6060",nil)
}
