package main

import (
	"awesomeProject/pc-client/user-portal"
	"fmt"
)

const (
	defaultName = "---world"
)

func main() {
	fmt.Println("client start.....")

	/*var invoker kit.KitInvoker
	invoker.InitInvoker()
	defer invoker.Close()

	c := hello.NewGreetClient(invoker.Conn)
	name := defaultName
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)*/

	/*u := pb.NewUserClient(invoker.Conn)
	rsp, err := u.Login(context.Background(), &pb.LoginRequest{Name: "client name"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("user login: %s", rsp.Message)*/
	//hello := hello_portal.NewHelloPortal()
	//hello.SayHelloTest("This is a test")
	//defer hello.Close()


	user := user_portal.NewUserPortal()
	user.UserLogin()
	defer user.Close()

}
