package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计英文，数字，空格和其他字符的个数
type CharCount struct {
	ChCount, NumCount, SpaceCount, OtherCount int
}

func main() {
	filePath := "e:/faling.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open File err", err)
		return
	}

	defer file.Close() //延迟关闭

	count := CharCount{}

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		//如果读道文件末尾
		if err == io.EOF {
			break
		}
		//遍历str
		for _, val := range str {
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough //穿透处理
			case val >= 'A' && val <= 'Z':
				count.ChCount++
			case val == ' ' || val == '\t':
				count.SpaceCount++
			case val >= '0' && val <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("统计结果：英文字母%v个,数字%v个,空格%v个,其他字符%v个\n", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
