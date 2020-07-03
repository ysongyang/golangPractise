package utils

import (
	"fmt"
	"math/rand"
	"time"
)

//生成订单号
//202006301593522964897081
func GenerateCode() string {
	date := time.Now().Format("20060102")
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, time.Now().UnixNano()/1e6, r)
	return code
}
