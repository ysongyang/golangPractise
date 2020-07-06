package main

import (
	"bufio"
	"fmt"
	"golangPractise/utils"
	"io"
	"io/ioutil"
	"os"
)

//打开一个文件，不存在则创建
func createFile() {
	file, err := os.OpenFile("e:/demo3.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()
	str := "hello,Gardong\r\n"
	//写入时带缓存的*Writer
	w := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		w.WriteString(str)
	}

	//writer 是带缓存的，因此调用writeString方法时，其实内容是先写入缓存，调用Flush方法时，将缓冲的数据写入到文件中
	w.Flush()
}

//覆盖文件内容
func coverFile() {
	file, err := os.OpenFile("e:/demo3.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()
	str := "你好，广东珠海\r\n"
	//写入时带缓存的*Writer
	w := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		w.WriteString(str)
	}

	//writer 是带缓存的，因此调用writeString方法时，其实内容是先写入缓存，调用Flush方法时，将缓冲的数据写入到文件中
	w.Flush()
}

//文件内容追加
func appendFile() {
	file, err := os.OpenFile("e:/demo3.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()
	str := "我是来追加内容的，其实内容是先写入缓存，调用Flush方法时，将缓冲的数据写入到文件中\r\n"
	//写入时带缓存的*Writer
	w := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		w.WriteString(str)
	}

	//writer 是带缓存的，因此调用writeString方法时，其实内容是先写入缓存，调用Flush方法时，将缓冲的数据写入到文件中
	w.Flush()
}

//O_RDWR  读写模式
func readFile() {
	file, err := os.OpenFile("e:/demo3.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()

	//先读
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行符结束一次
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "hello,珠海\r\n"
	//写入时带缓存的*Writer
	w := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		w.WriteString(str)
	}

	//writer 是带缓存的，因此调用writeString方法时，其实内容是先写入缓存，调用Flush方法时，将缓冲的数据写入到文件中
	w.Flush()
}

func copyFile() {
	//先将 e:/demo3.txt 读取到内存
	//在将文件写入 e:/demo3_copy.txt
	filePath1 := "e:/demo3.txt"
	filePath2 := "e:/demo3_copy.txt"

	if data, err := ioutil.ReadFile(filePath1); err != nil {
		fmt.Println("ReadFile err :", err)
		return
	} else {
		if err := ioutil.WriteFile(filePath2, data, 0666); err != nil {
			fmt.Println("WriteFile err :", err)
			return
		}
	}

}

func main() {

	/*createFile()
	coverFile()
	appendFile()*/
	//readFile()
	//copyFile()
	_, err := utils.CopyFile("d:/WIN10美化包.rar", "e:/WIN10美化包.rar")
	if err == nil {
		fmt.Println("copy ok")
	} else {
		fmt.Println(err)
	}
}
