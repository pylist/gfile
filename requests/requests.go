package requests

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

//type url map[string]interface{}

func Parse(params map[string]string) string {
	value := url.Values{}
	for k, v := range params {
		value.Add(k, v)
	}
	return value.Encode()
}

func Get(urls string, headers map[string]string, params map[string]string) ([]byte, error) {
	urls = urls + "?" + Parse(params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	//req.Header.Set("cookie", "COOKIE_LOGIN_USER=FF5F33F849B1AD7407AC19FEF689942FAE8C751D7A035BC9D5C0F268C8045F4AF2E2F5FE0D3677899278459B6EA59D91D2E983B83E652913B9B44B55")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}