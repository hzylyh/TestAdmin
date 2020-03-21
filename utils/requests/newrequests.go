package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"salotto/model/InterfaceTestPartEntity"
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

func (rq *Requests) Post(url, reqJson string, headerInfos []*InterfaceTestPartEntity.TInterfaceHeadersInfo) []byte {
	var err error
	rq.Requests, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJson)))
	//http.Post()

	for _, headerInfo := range headerInfos {
		rq.Requests.Header.Set("content-type", "application/x-www-form-urlencoded")
		rq.Requests.Header.Set(headerInfo.HeaderName, headerInfo.HeaderValue)
	}

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
