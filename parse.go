package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

//	This file contains operations to strings

// GetInfo ...
func GetInfo(str string) (string, error) {
	matched, err := regexp.MatchString(".*oldInfo:.*", str)
	infoRegexp := regexp.MustCompile(`.*oldInfo:\s(\{.*\})`)
	params := infoRegexp.FindStringSubmatch(str)
	if matched {
		return params[1], nil
	}
	return "", err
}

// GetStuNum ...
func GetStuNum(str string) (string, error) {
	matched, err := regexp.MatchString(".*number:.*", str)
	infoRegexp := regexp.MustCompile(`.*number:\s'(.*)'`)
	params := infoRegexp.FindStringSubmatch(str)
	if matched {
		return params[1], nil
	}
	return "", err
}

// GetMap 将字符串序列化为map，类型断言
func GetMap(str string) (map[string]string, error) {
	var data map[string]string
	// var info map[string]string
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return nil, err
	}
	return data, nil
}

// string好烦啊

// GetUser 读取用户信息
// func GetUser(str string) (map[string]string, error) {
// 	var data map[string]string
// 	// var info map[string]string
// 	if err := json.Unmarshal([]byte(str), &data); err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// FakeInfo 指定用户名，修改日期信息
func FakeInfo(username string) error {
	str := ReadString(username)
	var data map[string]interface{}
	// var info map[string]string
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return err
	}
	data["date"] = GetYMD()
	fmt.Println(data["date"])
	data["created"] = "1586920006"
	data["sfygtjzzfj"] = "0"
	data["sfsqhzjkk"] = "0"
	data["created_uid"] = "0"
	data["id"] = strconv.Itoa(300000 + rand.Intn(9)*10000 + rand.Intn(10000))
	byte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	str = string(byte)
	SaveString(str, username)
	return nil
}
