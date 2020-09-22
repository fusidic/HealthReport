package main

// NewRequestInfo init with login URL.
func NewRequestInfo(username, password string) RequestInfo {
	return RequestInfo{
		URL: "https://m.nuaa.edu.cn/uc/wap/login/check",
		Data: map[string]string{
			"username": username,
			"password": password,
		},
	}
}

// NewGetIndex ...
func NewGetIndex() GetIndex {
	return GetIndex{
		URL: "https://m.nuaa.edu.cn/ncov/wap/default/index",
	}
}

// NewRequestData init with health report information.
func NewRequestData(filename string) RequestData {
	data, _ := GetMap(ReadString(filename))
	return RequestData{
		URL:  "https://m.nuaa.edu.cn/ncov/wap/default/save",
		Data: data,
	}
}
