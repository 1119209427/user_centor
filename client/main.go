package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"user_centor/pb"
)

func main(){
	//链接服务
	grpcConn,err:=grpc.Dial("127.0.0.1:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatal(err.Error())
		return
	}
	defer grpcConn.Close()

	//起一个客户端
	grpcClient:=pb.NewUserServiceClient(grpcConn)
	var req pb.UserRequest
	req.Username="zhen xi"
	req.Password="123456"
	req.PasswordConfirm="123456"


	//调用远程函数
	resp,err:=grpcClient.UserRegister(context.TODO(),&req)
	fmt.Println(resp,err)


}
