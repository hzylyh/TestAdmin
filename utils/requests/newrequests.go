package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

type Requests struct {
	Requests *http.Request
	Client   *http.Client
}

func NewRequests() *Requests {
	jar, err := cookiejar.New(nil)

	if err != nil {
		panic(err)
	}

	return &Requests{
		Client: &http.Client{
			Jar: jar,
		},
	}
}

func (rq *Requests) Post(url, reqJson string) []byte {
	var err error
	rq.Requests, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJson)))
	//http.Post()
	rq.Requests.Header.Set("Content-Type", "application/json")

	resp, err := rq.Client.Do(rq.Requests)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("status", resp.Status)
	//fmt.Println("response:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	return body
}
