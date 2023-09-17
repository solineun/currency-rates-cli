package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL string = "https://www.cbr.ru/scripts/XML_daily.asp?date_req="
//date_req is dd/mm/yyyy format

func getXml(d Date) ([]byte, error) {
	url := concatUrlWithDate(d)
	fmt.Println(url)
	respBody, err := makeRequest(url)
	fmt.Println(string(respBody))
	return respBody, err
}

func concatUrlWithDate(d Date) string {
	return URL + d.cbrString()
}

func makeRequest(url string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Host", "www.cbr.ru")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	} 
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
}