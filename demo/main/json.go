package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

//结构体序列号案例
func jsonStruct() {
	monster := Monster{"张三", 30, "1989-05-05", 9290, "PHP,JAVA,GOLANG,PYTHON"}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))
}

func jsonMap() {
	var mp map[string]interface{}
	mp = make(map[string]interface{})
	mp["name"] = "王五"
	mp["age"] = 30
	mp["address"] = "广东 珠海"

	data, err := json.Marshal(mp)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))
}

func jsonSlice() {
	var slice []map[string]interface{}

	var m1, m2 map[string]interface{}

	m1 = make(map[string]interface{})

	m1["name"] = "王五"
	m1["age"] = 30
	m1["address"] = "广东 珠海"

	m2 = make(map[string]interface{})
	m2["name"] = "梨瓜"
	m2["age"] = 22
	m2["address"] = "河南 郑州"
	slice = append(slice, m1, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败...", err)
	}
	fmt.Println(string(data))
}

func main() {
	jsonStruct()
	jsonMap()
	jsonSlice()
}
