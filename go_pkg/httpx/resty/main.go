package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-resty/resty/v2"
)

var wg sync.WaitGroup

func main() {
	client := resty.New()

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go httpGet(i, client)
		// httpGetProm(i, client)
	}
	wg.Wait()

}

type Value struct {
	TodayNum    string `json:"todayNum"`
	YestodayNum string `json:"yestodayNum"`
}
type RES struct {
	Status string
	Data   Value
}

func httpGet(i int, client *resty.Client) {
	resValue := &RES{}
	// url := "http://192.168.10.85:8081/bc/api/monitor/response_time_with_date?requestDate=2022-01-04"
	url := "http://192.168.24.212:8080/bc/api/monitor/get_apdex?requestDate=2022-01-04"

	_, err := client.R().SetResult(resValue).Get(url)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Response Info:")
	// fmt.Println("Status Code:", resp.StatusCode())
	// fmt.Println(">>>>>>>", i)
	fmt.Printf("TodayNum:%s YestodayNum:%s\n", resValue.Data.TodayNum, resValue.Data.YestodayNum)
	wg.Done()
}

func httpGetProm(i int, client *resty.Client) {
	resValue := &ResTotal{}
	url0 := "http://192.168.10.85:49090/api/v1/query?query=avg%28easegress_api_response_time_with_date%7BrequestDate%3D%272022-01-03%27%7D%29&time=1641288269.214"

	url1 := "http://192.168.10.85:49090/api/v1/query?query=avg%28easegress_api_response_time_with_date%7BrequestDate%3D%272022-01-04%27%7D%29&time=1641288269.214"

	_, err := client.R().SetResult(resValue).Get(url0)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("yestoday", resValue.Data.Result)
	_, err = client.R().SetResult(resValue).Get(url1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("today:", resValue.Data.Result)
	wg.Done()
}

type ResTotal struct {
	Status string      `json:"status"`
	Data   DataStruct1 `json:"data"`
}
type DataStruct1 struct {
	ResultType string       `json:"resultType"`
	Result     []ResStruct1 `json:"result"`
}
type ResStruct1 struct {
	Metric interface{}   `json:"metric"`
	Value  []interface{} `json:"value"`
}
