package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string = "https://www.cbr.ru/scripts/XML_daily.asp?date_req="
//date_req = Date of query (dd/mm/yyyy)

func getXml(d Date) ([]byte, error) {
	url = concatUrlWithDate(d)
	respBody, err := makeRequest(url)
	return respBody, err
}

func concatUrlWithDate(d Date) string {
	return url + fmt.Sprintf("%d/%d/%d", d.Day, d.Month, d.Year)
}

func makeRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
}