package main

import (
	"fmt"
	"net"
)


func main() {
	
	fmt.Println("客户端已启动~正在监听广播~");

	listener,listenErr := net.Listen("tcp",":8001");
	
	check("listenErr",listenErr);

	defer listener.Close();

	conn,AcceptErr := listener.Accept();

	check("AcceptErr",AcceptErr);

	defer conn.Close();		
	
	// 监听
	for{

		// 新建一个切片存放收听到的
		can := make([]byte,255);

		_,readErr := conn.Read(can);

		check("readErr",readErr);

		conn.Write(can);

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