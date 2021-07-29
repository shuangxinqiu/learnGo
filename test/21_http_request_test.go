package test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"testing"
)

/*
获取GET参数
方法列表
方法名												描述
(r *Request) ParseForm() error						判断是否解析传参时出错
(r *Request) FormValue(key string) string			接收指定key的参数值
*/
//服务端代码
//curl -i 'localhost:8888/login?userName=admin&password=1234'
func TestHttpReqGet(t *testing.T) {
	http.HandleFunc("/login", login2)
	_ = http.ListenAndServe(":8888", nil)
}
func login2(w http.ResponseWriter, r *http.Request) {
	//判断参数是否Get请求，并且参数解析正常
	if r.Method == "GET" && r.ParseForm() == nil {
		//接收参数
		userName := r.FormValue("userName")
		fmt.Printf("userName: %s \n", userName)
		password := r.FormValue("password")
		fmt.Printf("password: %s \n", password)
		if userName == "" || password == "" {
			w.Write([]byte("用户名或密码不能为空"))
		}
		if userName == "admin" && password == "1234" {
			w.Write([]byte("登录成功"))
		} else {
			w.Write([]byte("用户名或密码错误"))
		}

	}
}

/*
获取POST参数
接收POST参数分以下两种情况
1.普通的Post表单：Content-Type=application/x-www-form-urlencoded
2.有文件上传的Post表单：Content-Type=multipart/form-data
*/
//普通的Post表单（r.PostFormValue)
//curl -i -d 'userName=admin&password=1234' 'localhost:8888/login'
func TestPostFormValue(t *testing.T) {
	http.HandleFunc("/login", login3)
	_ = http.ListenAndServe(":8888", nil)
}
func login3(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.ParseForm() == nil {
		//接收post参数
		userName := r.PostFormValue("userName")
		fmt.Printf("userName: %s \n", userName)
		password := r.PostFormValue("password")
		fmt.Printf("password: %s \n", password)
		if userName == "" || password == "" {
			w.Write([]byte("用户名或密码不能为空"))
			return
		}
		if userName == "admin" && password == "1234" {
			w.Write([]byte("登录成功！"))
			return
		}
		w.Write([]byte("用户名或密码错误"))
	} else {
		w.Write([]byte("当前接口，仅支持POST请求！"))
	}
	return
}

//有文件上传的Post表单（r.FormFile)
func TestFormFile(t *testing.T) {
	http.HandleFunc("/upload", upload)
	_ = http.ListenAndServe(":8888", nil)
}

//上传文件
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.ParseForm() == nil {
		//接收上传文件的参数
		formFile, fileHeader, err := r.FormFile("file")
		if formFile == nil {
			http.Error(w, "上传文件不能为空！", http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//延迟关闭文件
		defer formFile.Close()
		//获取上传文件的名称
		fileName := fileHeader.Filename
		fmt.Printf("文件名称：%s \n", fileName)
		//获取文件的后缀
		ext := path.Ext(fileName)
		fmt.Printf("文件后缀：%s \n", ext)
		//创建新文件，如果同名文件存在，则会清空
		pathName := "public/img"
		if !pathExist(pathName) {
			err := os.MkdirAll(pathName, os.ModePerm)
			if err != nil {
				http.Error(w, "创建目录失败"+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		newFile, err := os.Create(pathName + "/" + fileName)
		if err != nil {
			http.Error(w, "创建文件失败："+err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()
		//将formFile复制到newFile，从而实现上传的功能
		written, err := io.Copy(newFile, formFile)
		fmt.Printf("上传结果：%d \n", written)
		if err != nil {
			http.Error(w, "文件上传失败："+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("文件上传成功！"))
	}
	return
}

//判断目录是否存在
func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}


/*
获取Cookie中的值
cookie的数据结构
Cookie是一个结构体，其中有Cookie的名和值、domain、过期时间等信息，具体定义方式：
type Cookie struct {
	Name string //变量名
	Value string //变量值
	Path string //设置访问哪些路径，应该携带这个cookie，不设置则代表所有
	Domain string //设置访问域名范围的主机时，应该携带这个cookie，不设置则代表所有
	Expires time.Time //一个时间值，代表什么时候过期
	RawExpires string //for reading cookies only
	MaxAge int //用来设置过期，为负数或等于0表示立即过期，大于0表示多少秒之后过期
	Secure bool
	HttpOnly bool
	SameSite SameSite
	Raw string
	Unparsed []string //Raw text of unparsed attribute-value pairs
}
*/