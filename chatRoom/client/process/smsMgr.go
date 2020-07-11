package process

import (
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
)

func outputGroupMsg(mes *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(mes.Data), &smsMsg)
	if err != nil {
		fmt.Println("outputGroupMsg json error", err.Error())
		return
	}

	//显示内容格式化输出
	info := fmt.Sprintf("用户ID【%d】对大家说：%s\n", smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
	fmt.Println()
}
