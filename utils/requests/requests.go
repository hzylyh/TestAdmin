package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"salotto/utils"
	"salotto/utils/cronUtil"
)

func main() {
	//url := "http://localhost:8089/salotto/itfPart/test"
	//reqJson := "{\"UserId\":\"" + usrId + "\",\"Password\":\"" + pwd + "\"}"
	//reqJson := "{\"UserId\":\"123\",\"Password\":\"4567\"}"
	//Post(url, reqJson)

	var exp, act string
	row := cronUtil.CronDB.Table("test_assert").Where("id = ?", 1).Select("exp, act").Row()
	row.Scan(&exp, &act)

	verifyCol := []string{"resultEntity.isRead", "resultEntity.tCardList"}

	utils.MulAssert(exp, act, verifyCol)
}

func Post(url, reqJson string) {
	//url = "http://www.baidu.com"

	//json序列化
	//post := "{\"UserId\":\"" + usrId + "\",\"Password\":\"" + pwd + "\"}"

	fmt.Println(url, "post", reqJson)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqJson)))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
