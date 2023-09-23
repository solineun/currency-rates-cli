package model

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/golang-module/carbon/v2"
)

const URL string = "https://www.cbr.ru/scripts/XML_daily.asp?date_req="
//date_req is dd/mm/yyyy format

func getXml(d Date) ([]byte, error) {
	url := concatUrlWithDate(d)
	respBody, err := makeRequest(url)
	return respBody, err
}

func concatUrlWithDate(d Date) string {
	return URL + d.cbrString()
}

func makeRequest(cbrUrl string) ([]byte, error) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest(
		"GET", cbrUrl, nil,
	)
	if err != nil {
		return nil, err
	}

	urlObj, _ := url.Parse(cbrUrl)
	cooks := getCookies()
	client.Jar.SetCookies(urlObj, cooks)

	addHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} 
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
}

func addHeaders(req *http.Request) {
	req.Header.Add("Host", "www.cbr.ru")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
}

func getCookies() []*http.Cookie {
	return []*http.Cookie{
		{
			Name: "__ddg1_",
			Path: "/",
			Domain: "cbr.ru",
			HttpOnly: true,
			Expires: carbon.Now().AddYear().Carbon2Time(),
		},
		{
			Name: "accept",
			Value: "1",
		},
	}
}