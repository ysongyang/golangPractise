package testCaseExce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type monster struct {
	Name  string
	Age   int
	Skill string
}

const filePath = "e:/monster.log" //定义日志文件

//绑定一个序列化方法
func (m *monster) Store() bool {
	//序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return false
	}
	//保存到文件

	if err := ioutil.WriteFile(filePath, data, 0666); err != nil {
		fmt.Println("写文件错误，WriteFile：", err)
		return false
	}
	return true
}

//绑定一个方法,反序列化
func (m *monster) Restore() bool {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("反序列化失败：", err)
		return false
	}

	//读取到的data []byte ，反序列化
	if err := json.Unmarshal(data, m); err != nil {
		fmt.Println("反序列化错误，Unmarshal：", err)
		return false
	}
	return true
}


