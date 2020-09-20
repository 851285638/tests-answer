package main

// server.go

import (
	"crypto/rand"
	"fmt"
	pb "goMicroService/msgConfig"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math"
	"math/big"
	"net"
	"strconv"
	"sync"
)

const (
	port = ":13014"
)

type server struct {}

var users map[string]pb.UserDetail
var ids map[int64]int

func (s *server) Login(ctx context.Context, login *pb.LoginRequest) (*pb.LoginReply, error) {
	userName := login.Name
	password := login.Password
	if user ,ok := users[userName]; ok {
		if user.Password == password {
			//返回值为1说明登录成功，同时返回用户id，返回值为0说明用户名或密码错误
			//这块可以添加检测用户在线状态以及对应处理的逻辑
			fmt.Println("Here is a user login successfully !")
			return &pb.LoginReply{Id: strconv.FormatInt(user.Id, 10), State: 1}, nil
		}
		fmt.Println("Here is a new user login failure, because password wrong !")
		return &pb.LoginReply{State: 0}, nil
	}
	fmt.Println("Here is a new user login failure, because account wrong !")
	return &pb.LoginReply{State: 0}, nil
}
func (s *server) Register(ctx context.Context, register *pb.RegisterRequest) (*pb.RegisterReply, error) {
	var l sync.Mutex
	userName := register.Detail.Account
	password := register.Detail.Password
	if _ ,ok := users[userName]; !ok {
		age := register.Detail.Age
		id := getNewUserId()
		l.Lock()
		users[userName] = pb.UserDetail{Account: userName, Password: password, Age: age, Id: id}
		l.Unlock()
		//返回值为1说明注册成功
		fmt.Println("Here is a new user registing !")
		return &pb.RegisterReply{State: 1}, nil
	}else {
		//返回值为0说明用户名重复
		fmt.Println("A new user registed failure, because account already existed !")
		return &pb.RegisterReply{State: 0}, nil
	}
}

func main() {
	ids = make(map[int64]int, 0)
	users = make(map[string]pb.UserDetail, 0)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	s.Serve(lis)
}

func getNewUserId() int64 {
	var id int64
	var l sync.Mutex
	for {
		n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if _ ,ok := ids[n.Int64()]; !ok {
			l.Lock()
			ids[n.Int64()] = 1
			l.Unlock()
			id = n.Int64()
			break
		}
	}
	return id
}