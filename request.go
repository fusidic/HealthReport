package main

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// RequestInfo 保存Login网址及信息
type RequestInfo struct {
	URL  string            //post对应的网址
	Data map[string]string //post要传输的数据，必须key value必须都是string
}

// Login 适用于 application/x-www-form-urlencoded
func (r RequestInfo) Login(client *http.Client) ([]byte, error) {

	//post要提交的数据
	DataURLVal := url.Values{}
	for k, v := range r.Data {
		DataURLVal.Add(k, v)
	}
	req, err := http.NewRequest("POST", r.URL, strings.NewReader(DataURLVal.Encode()))
	if err != nil {
		return nil, err
	}

	//伪装头部
	req.Header.Set("Host", "m.nuaa.edu.cn")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "34")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("DNT", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://m.nuaa.edu.cn")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://m.nuaa.edu.cn/uc/wap/login?redirect=https%3A%2F%2Fm.nuaa.edu.cn%2Fncov%2Fwap%2Fdefault%2Findex")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	// we don't need Cookie here
	// req.Header.Add("Cookie", "fc_session=eyJpdiI6InRjN1l3akRxSTBTc2NRcWd5YUtoQ2c9PSIsInZhbHVlIjoiZ0dmTHkwTnpON21nUVlKTmlvbjdUY0U5NUFtbW1qZ1dteXByd0JxTmJqTUVBU0RiVmVYZURncXFrWG5ucFlSWmwwRmV6MkNVZEwyQVwvbXhWRnRWSldRPT0iLCJtYWMiOiI4ZDZmYTVlM2I4NWNkMzY2YjY0YTExOTVhNzUzYTFhZWE5NGZmNjIyNzVhZTA0YWE1YTQ2OGE4NTMwZjcyNzQyIn0%3D; old_device_token=f5ad540923c1512e02a9769d2d30a7c6; device_token=f2038259c310e236f1eb4805fb03272d; Hm_lvt_48b682d4885d22a90111e46b972e3268=1584947651; eai-sess=i37mofdnh8hdr0t15qlov06r84; UUkey=307670b833493ccdd678b5ed3497b682; Hm_lpvt_48b682d4885d22a90111e46b972e3268=1584979394")
	// req.Header.Add("Cookie", "fc_session=eyJpdiI6InRjN1l3akRxSTBTc2NRcWd5YUtoQ2c9PSIsInZhbHVlIjoiZ0dmTHkwTnpON21nUVlKTmlvbjdUY0U5NUFtbW1qZ1dteXByd0JxTmJqTUVBU0RiVmVYZURncXFrWG5ucFlSWmwwRmV6MkNVZEwyQVwvbXhWRnRWSldRPT0iLCJtYWMiOiI4ZDZmYTVlM2I4NWNkMzY2YjY0YTExOTVhNzUzYTFhZWE5NGZmNjIyNzVhZTA0YWE1YTQ2OGE4NTMwZjcyNzQyIn0%3D; old_device_token=f5ad540923c1512e02a9769d2d30a7c6; device_token=f2038259c310e236f1eb4805fb03272d; Hm_lvt_48b682d4885d22a90111e46b972e3268=1584947651; eai-sess=i37mofdnh8hdr0t15qlov06r84; UUkey=307670b833493ccdd678b5ed3497b682; Hm_lpvt_48b682d4885d22a90111e46b972e3268=1585014478")

	//提交请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Printf("resp: %+v\n\n", resp)

	//读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result)

	return result, nil
}

// GetIndex 获取相关信息
type GetIndex struct {
	URL string
}

// GetReport 从index页面中获取一些信息
func (g GetIndex) GetReport(client *http.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", g.URL, nil)
	if err != nil {
		return nil, err
	}

	//伪装头部
	req.Header.Set("Host", "m.nuaa.edu.cn")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Control", "max-age=0")
	req.Header.Set("DNT", "1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Referer", "https://m.nuaa.edu.cn/uc/wap/login?redirect=https%3A%2F%2Fm.nuaa.edu.cn%2Fncov%2Fwap%2Fdefault%2Findex")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	//提交请求
	response, err := client.Do(req)

	//解决Content乱码问题
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		defer reader.Close()
	default:
		reader = response.Body
	}

	// io.Copy(os., reader) // print html to standard out
	result, err := ioutil.ReadAll(reader)
	return result, nil
}

// RequestData 保存上报网址及信息
type RequestData struct {
	URL  string
	Data map[string]string
}

func (r RequestData) postSave(client *http.Client) ([]byte, error) {
	//post要提交的数据
	DataURLVal := url.Values{}
	for key, val := range r.Data {
		DataURLVal.Add(key, val)
	}

	req, err := http.NewRequest("POST", r.URL, strings.NewReader(DataURLVal.Encode()))
	if err != nil {
		return nil, err
	}
	//伪装头部
	req.Header.Set("Host", "m.nuaa.edu.cn")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "2015")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("DNT", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://m.nuaa.edu.cn")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://m.nuaa.edu.cn/ncov/wap/default/index")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	//提交请求
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	//读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result)
	return result, nil
}
