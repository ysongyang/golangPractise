//工具类
package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
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

//判断当前文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

//文件拷贝
func CopyFile(dst string, src string) (written int64, err error) {

	flag, err := PathExists(src)

	if !flag {
		return 0, errors.New("srcFile is not exists")
	}

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Open srcFile err", err)
	}
	//关闭文件句柄
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("openFile dstFile err", err)
	}
	//关闭文件句柄
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}
