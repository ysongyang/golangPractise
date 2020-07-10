package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golangPractise/chatRoom/common/message"
	"net"
)

//方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [1024 * 4]byte //传输数据的缓存区大小
}

//读数据
func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {

	//读取4个字节，先读取长度Unmarshal
	//这里会造成阻塞
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		return
	}
	//数据buf[:4] 转成一个 uint32类型
	pkgLen := binary.BigEndian.Uint32(transfer.Buf[:4])
	//根据pakLen 读取消息内容  (从 conn套接字里读取pkgLen长度的数据放入buf中去）
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}
	//pkgLen反序列化 -> message.Message
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes) //不加&符号，mes 可能为nil
	if err != nil {
		return mes, fmt.Errorf("readPkg json Unmarshal error: %s", err)
	}
	//fmt.Printf("收到客户端 %v 发来的内容长度为 %v\n", conn.RemoteAddr(), buf[:4])
	return mes, err
}

//发数据
func (transfer *Transfer) WritePkg(data []byte) (err error) {

	//1. 先获取长度 转成一个表示长度的切片
	pkgLen := uint32(len(data))
	//长度转成切片
	binary.BigEndian.PutUint32(transfer.Buf[0:4], pkgLen)
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		return fmt.Errorf("writePkg conn.Write buf error: %s", err)
	}
	//fmt.Printf("客户端发送数据消息长度=%d 内容=%s\n", len(dataMsg), string(dataMsg))

	//发送消息体
	n, err = transfer.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		return fmt.Errorf("writePkg conn.Write data error: %s", err)
	}
	return
}
