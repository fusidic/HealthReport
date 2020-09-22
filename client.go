package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

// DoClient finish report for one user{"username":"password"}
func DoClient(username, password string) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	// 构建/login/check请求
	req1 := NewRequestInfo(username, password)
	res1, err := req1.Login(client)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(res1))

	// 判断是否已存在username.json
	path := "./info/" + username + ".json"
	if !fileExists(path) {
		//获取主页信息
		report := NewGetIndex()
		res2, err2 := report.GetReport(client)
		if err2 != nil {
			fmt.Println(err)
		}
		// 从中解析oldInfo的值，并保存到本地
		oldInfo, _ := GetInfo(string(res2))
		SaveString(oldInfo, username)
	}

	// 更新oldInfo的信息
	errF := FakeInfo(username)
	if errF != nil {
		fmt.Println(errF)
	}

	req2 := NewRequestData(username)
	// fmt.Println(req2)
	resp, err := req2.postSave(client)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(resp))
	fmt.Printf("User: %s finished\n", username)
	fmt.Printf("Login msg: %s \nHealth report: %s \n\n", string(res1), string(resp))
	// modify SendMail in mail.go to complete this function
	// SendMail(str string)
}
