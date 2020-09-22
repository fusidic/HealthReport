package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// SaveString 保存个人信息到对应的学号中
func SaveString(str string, filename string) error {
	var path string
	path = "./info/" + filename + ".json"
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	l, err := f.WriteString(str)
	if err != nil {
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	return err
}

// ReadString (filename string) 从本地读取filename信息，不要后缀名
func ReadString(filename string) string {
	var path string
	path = "./info/" + filename + ".json"
	if !fileExists(path) {
		return ""
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	return text
}
