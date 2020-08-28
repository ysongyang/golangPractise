package main

import (
	"encoding/json"
	"fmt"
	"golangPractise/webApp/model"
	"net/http"
)

/*
type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	fmt.Fprintln(w, "自定义的ServeHttp", r.URL.Path)
}
*/

type Result struct {
	Code    int           `json:"code"`
	Data    []*model.User `json:"data"`
	Message string        `json:"message"`
}

//创建处理器
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	w.Header().Add("server", "golang-http-server")
	w.Header().Add("status", "200")
	//设置请求类型
	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintln(w, "发送的请求地址是：", r.URL.Path)
	//fmt.Fprintln(w, "发送的请求地址后的查询字符串是：", r.URL.RawQuery)
	//fmt.Fprintln(w, "发送的请求头中的所有信息有：", r.Header)
	//fmt.Fprintln(w, "发送的请求头中 User-Agent的信息有：", r.Header["User-Agent"])
	//fmt.Fprintln(w, "发送的请求头中 User-Agent的属性值：", r.Header.Get("User-Agent"))
	//解析表单 需要放在前面
	//r.ParseForm()

	if r.Method == "GET" {
		//在r.Form 之前需要先调用ParseForm()
		//fmt.Fprintln(w, "获取GET请求参数有：", r.Form)

		//fmt.Fprintln(w, "获取GET请求参数username：", r.FormValue("username"))
		//fmt.Fprintln(w, "获取GET请求参数password：", r.FormValue("password"))
	} else if r.Method == "POST" {

		//这里r.Body.Read 把数据内容转换成切片，所有r.ParseForm解析后可能无法获取到请求体内容  string(body) 可能为乱码
		//获取请求体内容长度
		/*len := r.ContentLength
		//创建一个切片并make
		body := make([]byte, len)
		//将请求Body里的内容读取到body
		r.Body.Read(body)
		//请求方式 x-www-form-urlencoded
		fmt.Fprintln(w, "获取Post请求体的内容：", string(body))*/

		//在r.PostForm 之前需要先调用ParseForm()
		//fmt.Fprintln(w, "获取Post请求参数有：", r.PostForm)

		//fmt.Fprintln(w, "获取Post请求参数username：", r.PostFormValue("username"))
		//fmt.Fprintln(w, "获取Post请求参数password：", r.PostFormValue("password"))

	}
	user := &model.User{}
	users, _ := user.GetUsers()
	ret := &Result{
		Code:    0,
		Message: "Success",
		Data:    users,
	}
	json, _ := json.Marshal(ret)
	w.Write(json)

}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	//设置状态码
	w.WriteHeader(301)
}

func main() {
	fmt.Println("Golang Http Server Start ...")
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/redirect", redirect)

	/**自定义Handler**/
	//myHandler := MyHandler{}
	//http.Handle("/myHandler", &myHandler)
	//server := http.Server{
	//	Addr:    ":8080",
	//	Handler: &myHandler,
	//}
	//server.ListenAndServe()
	/**自定义Handler**/

	//创建路由
	http.ListenAndServe(":8080", nil)
}
