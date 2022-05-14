package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_centor/pb"
	"user_centor/service/core"
)

func main(){
	//初始化一个grpc
	grpcServer:=grpc.NewServer()
	//注册服务
	pb.RegisterUserServiceServer(grpcServer,new(core.UserServices))
	//设置监听，指定port和ip
	listener,err:=net.Listen("tcp","127.0.0.1:8080")
	if err!=nil{
		fmt.Println("监听出错")
		log.Fatal(err.Error())
		return
	}
	defer listener.Close()
	//启动服务
	fmt.Println("服务启动")
	grpcServer.Serve(listener)
}
