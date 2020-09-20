package main

//client.go

import (
	pb "goMicroService/msgConfig"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:13014"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTestClient(conn)
	//正常注册测试
	r, err := c.Register(context.Background(), &pb.RegisterRequest{Detail: &pb.UserDetail{Account: "123", Password: "123", Age: 24}})
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	log.Printf("State: %d", r.State)

	//重复用户名注册测试
	r1, err1 := c.Register(context.Background(), &pb.RegisterRequest{Detail: &pb.UserDetail{Account: "123", Password: "123", Age: 24}})
	if err1 != nil {
		log.Fatal("Error: %v", err)
	}
	log.Printf("State: %d", r1.State)

	//正确信息登录测试
	r2, err2 := c.Login(context.Background(), &pb.LoginRequest{Name: "123", Password: "123"})
	if err2 != nil {
		log.Fatal("Error: %v", err)
	}
	log.Printf("State: %d, id : %s", r2.State, r2.Id)

	//错误信息登录测试
	r3, err3 := c.Login(context.Background(), &pb.LoginRequest{Name: "1234", Password: "123"})
	if err3 != nil {
		log.Fatal("Error: %v", err)
	}
	log.Printf("State: %d, id : %s", r3.State, r3.Id)
}