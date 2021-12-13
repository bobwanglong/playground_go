package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	url := "http://localhost:10099/hello/1"
	var wg sync.WaitGroup
	for i := 0; i < 150; i++ {
		if i%2 == 0 {
			wg.Add(1)
			go httpDelete(url, &wg)
			continue
		}

		wg.Add(1)
		go httpGet(url, &wg)
	}
	wg.Wait()

}

func httpGet(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(string(body))
	wg.Done()
}

func httpDelete(url string, wg *sync.WaitGroup) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(string(body))
	wg.Done()
}
