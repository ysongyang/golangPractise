package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

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
