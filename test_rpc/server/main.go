package  main

import (
"fmt"
"io"
"net"
"net/http"
"net/rpc"
)

type Watcher int


//该方法必须大写，第一个为接受的参数,第二个参数必须为指针为返回的结果,还必须返回值为error
func (w *Watcher) GetInfo(arg int, result *int) error {
	*result = 1
	return nil
}

func main() {

	http.HandleFunc("/ghj1976", Ghj1976Test)

	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败，端口可能已经被占用")
	}
	fmt.Println("正在监听1234端口")
	http.Serve(l, nil)
}

func Ghj1976Test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>ghj1976-123</body></html>")
}
