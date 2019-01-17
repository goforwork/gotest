package main

import (
	"fmt"
	"net"
)


func main() {

	var (
		sendStr string
	)

	fmt.Println("这里是服务端");
	
	// 创建客服端连接
	conn,dialDrr := net.Dial("tcp","127.0.0.1:8001");

	check("dialDrr",dialDrr);

	defer conn.Close();

	// 监听用户返回
	go listenClient(conn);

	// 循环监听键盘输入，将输入的数据发送给客户端
	for{

		fmt.Println("请输入广播内容：");

		fmt.Scanln(&sendStr);

		// 发送
		conn.Write([]byte(sendStr));

	}



}


// 新启一个协程监听客服端返回
func listenClient(conn *net.Conn) {

	for{

		can := make([]byte, 255);

		_,readErr := *conn.Read(can);

		fmt.Println(string(can));

	}

}


// 错误处理
func check(errType string, e error) {
    if e != nil {
    	fmt.Println(errType)
        panic(e)
    }
}