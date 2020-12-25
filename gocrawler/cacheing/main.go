package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	resp, err := http.Get("https://studygolang.com/")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	stringBody := string(data)
	count := strings.Count(stringBody, "<a")
	fmt.Println(count)

}