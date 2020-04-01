package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// URL in commandline example: http://www.gopl.io
func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	//resp, err := http.Get("http://www.gopl.io/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading")
		os.Exit(1)
	}
	fmt.Printf("%s", b)

}
