package test

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Go语言标准库内置的net/http包，可以实现HTTP服务端。实现HTTP服务端就是能够启动Web服务，相当于搭建起了一个Web服务器。
http.ListenAndServer()函数用来启动Web服务，绑定并监听http端口。

//addr：监听地址； 如 :8080 或者 0.0.0.0:8080            handler: HTTP处理器Handler
func ListenAndServer(addr string, handler Handler)
*/

/*
启动Web服务的几种方式
根据不同服务返回的handler，常见启动Web服务有以下几种方式。
*/
/*
http.FileServer:静态文件服务
http.FileServer()搭建的服务器只提供静态文件的访问。因为这种Web服务只支持静态文件访问，所以称之为静态文件服务。
*/
func TestHttp(t *testing.T) {
	runFileServer()
}
func runFileServer() {
	//如果该路径下有index.html，则会优先显示index.html，否则会看到文件目录
	http.ListenAndServe(":3000", http.FileServer(http.Dir("./public/")))
}

/*
http.HandlerFunc:默认的多路由分发服务
http.HandlerFunc()的作用是注册网络访问的路由。因为它采用的是默认的路由分发任务方式，所以称之为默认的多路由分发服务。

func HandlerFunc(pattern string, handler func(ResponseWriter, *Request){
	DefaultServeMux.HandlerFunc(patter, hanlder)
}
pattern:请求路径的匹配模式
handler:函数类型，表示这个请求需要处理的事情，其实就是Handler接口中ServeHTTP()方法。
ServeHTTP()方法有两个参数，其中第一个参数是ResponseWriter类型，包含了服务器端给客户端的响应数据。服务器端往ResponseWriter写了什么内容，浏览器的网页源码是什么内容。第二个参数一个*Request指针，包含了客户端发送给服务器的请求信息（路径、浏览器类型等）

通过http.HandlerFunc()注册网络路由时，http.ListenAndServe()的第二参数通常为nil，这意味着服务端采用默认的http.DefaultServeMux进行分发处理。
*/
func TestHttpHandlerFunc(t *testing.T) {
	//绑定路由hello
	http.HandleFunc("/hello", helloHandler)

	//绑定路由test
	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":5000", nil)
	fmt.Println("err: ", err)
}

//处理路由Hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("访问路由hello")
	//解析url参数
	fmt.Println(r.URL.Query())
	w.Write([]byte("hello world"))
}

//处理路由Test
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("访问路由test")
	//解析url参数
	fmt.Println(r.URL.Query())
	w.Write([]byte("test doing"))
}


/*
http.NewServeMux(): 自定义多路由分发服务
http.NewServeMux()的作用是注册网络访问的多路路由。因为它采用的是自定义的多路由分发任务方式，所以称之为自定义多路由分发服务。
注册网络路由时，如果http.ListenAndServer()的第二个参数为nil，那么表示服务端采用默认的http.DefaultServeMux进行分发处理。也可以自定义ServeMux。ServeMux结构体如下：
ServeMux结构体源码：
type ServeMux struct {
	mu sync.RWMutex //锁，由于请求涉及到并发处理，因此这里有个锁机制
	m	map[string]muxEntry //存放具体的路由信息
	es	[]muxEntry //按照路由长度从大到小的存放处理函数
	hosts  bool		//标记路由中是否带有主机名
}

//muEntry是路由的具体条目
type muxEntry struct {
	h Handler //处理函数
	pattern string //路由路径
}
*/
//自定义多路由实践
//定义一个接口体，用来实现http.Handler
type MyRoute struct {

}
//只实现http.Handler接口中的ServeHTTP方法
func (m *MyRoute) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	switch path {
	case "/":
		w.Write([]byte("首页"))
	case "/hello":
		w.Write([]byte("say hello"))
	case "/test":
		w.Write([]byte("say test"))
	default:
		http.NotFound(w,r)
	}
	return
}

func TestCustomHttpRoute(t *testing.T) {
	myRoute := &MyRoute{}
	http.ListenAndServe(":10000",myRoute)
}

/*
代码执行流程
使用http.ListenAndServe(":10000",myRoute)启动服务之后，会发生以下操作

1.实例化http.Server,并调用ListenAndServer()
func ListenAndServer(addr string, handler Handler) error {
	//实例化 Server
	server := &Server{Addr: addr, Handler: handler}
	//调用 ListenAndServe()
	return server.ListenAndServe()
}

2.监听端口
func (srv *Server) ListenAndServe() error {
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	//监听端口
	ln, err := net.Listen("tcp",addr)
	if err != nil {
		return err
	}
	//启动服务
	return srv.Serve(ln)
}

3.启动for无限循环，在循环体中Accept请求，并开启goroutine为这个请求服务
func (srv *Server) Serve(l net.Listener) error {
	//省略n行代码
	for{
		rw,err := l.Accept()
		if err != nil{
			//省略n行代码
		}
		connCtx := ctx
		if cc := srv.ConnContext; cc != nil{
			connCtx = cc(connCtx, rw)
			if connCtx == nil{
				panic("ConnContext returned nil")
			}
		}
	}
	tempDelay = 0
	c := srv.newConn(rw)
	c.setState(c.rwc, StateNew)
	//开启goroutine为这个请求服务
	go c.serve(connCtx)
}

4.读取每个请求内容，并调用ServeHTTP
func (c *conn) serve(ctx context.Context) {
 //...省略N行代码
 for {
    // 读取每个请求内容
  w, err := c.readRequest(ctx)
    //...省略N行代码

    // 调用ServeHTTP
  serverHandler{c.server}.ServeHTTP(w, w.req)
  //...省略N行代码
}

5. 判断handler是否为空,如果为空则把handler设置成DefaultServeMux。
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
 handler := sh.srv.Handler
 if handler == nil {
    // 如果为空则把handler设置成DefaultServeMux。
  handler = DefaultServeMux
 }
 if req.RequestURI == "*" && req.Method == "OPTIONS" {
  handler = globalOptionsHandler{}
 }
  // 上述示例中，传的是&MyRoute，所以会调用MyRoute.ServeHTTP
 handler.ServeHTTP(rw, req)
}
*/