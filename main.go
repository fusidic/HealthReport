package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	content := ReadString("user")
	var userdata []map[string]string
	if err := json.Unmarshal([]byte(content), &userdata); err != nil {
		fmt.Println(err)
	}
	for _, m := range userdata {
		username := m["username"]
		password := m["password"]
		DoClient(username, password)
	}
}
