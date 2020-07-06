//工具类
package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//邮箱校验
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//手机号校验
func VerifyMobileFormat(mobile string) bool {
	result, _ := regexp.MatchString(`^(1[3|4|5|7|8|9][0-9]\d{4,8})$`, mobile)
	if result {
		return true
	} else {
		return false
	}
}

//生成订单号
//202006301593522964897081
func GenerateCode() string {
	date := time.Now().Format("20060102")
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, time.Now().UnixNano()/1e6, r)
	return code
}
