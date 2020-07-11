package processes

import "fmt"

//服务端维护用户在线的map 类型是 UserProcess
//保证UserMgr 在服务端只有一个实例

var userMgr *UserMgr //全局变量

type UserMgr struct {
	onLineUsers map[int]*UserProcess
}

//初始化userMgr
func init() {
	userMgr = &UserMgr{
		onLineUsers: make(map[int]*UserProcess, 1024),
	}
}

//对userMgr的增删改查操作
func (um *UserMgr) AddOnlineUser(userProcess *UserProcess) {
	um.onLineUsers[userProcess.UserId] = userProcess
}

//删除
func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onLineUsers, userId)
}

//获取在线的用户
func (um *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return um.onLineUsers
}

//查询某个用户是否在线（根据id返回对应的process）
func (um *UserMgr) GetOnlineUserUserById(userId int) (up *UserProcess, err error) {
	userProcess, ok := um.onLineUsers[userId]
	if !ok { //说明当前用户是不在线
		err = fmt.Errorf("当前用户%d不在线\n", userId)
		return
	}
	return userProcess, nil
}
