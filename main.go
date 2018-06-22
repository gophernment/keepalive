package main

import (
	"fmt"
	"net/http"
)

func main() {
	ok, fail := 0, 0
	for i := 0; i < 6000; i++ {
		s, f := call(i)
		ok += s
		fail += f
	}

	fmt.Println("success", ok)
	fmt.Println("fail", fail)
}

func call(i int) (int, int) {

	url := "http://localhost:1323/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	c := http.DefaultClient
	// c := &http.Client{
	// 	Transport: &http.Transport{
	// 		DisableKeepAlives: true,
	// 	},
	// }

	_, err = c.Do(req)
	if err != nil {
		fmt.Println(i, err)
		return 0, 1
	}
	return 1, 0
}
