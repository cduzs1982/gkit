package main

import (
	pb "awesomeProject/proto/user"
	"awesomeProject/server/service"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()

	fmt.Println("---service start---"+dir)

	//s := kit.Init("C:\\Users\\AZeng\\go\\src\\awesomeProject\\server\\kit-config.yaml")
	//hello.RegisterGreetServer(s, &service.HelloService{})
	//

	//hel := service.NewHelloService()
	//hello.RegisterGreetServer(hel.RPC, hel)
	//hel.StartService()

	us := service.NewUserService()
	pb.RegisterUserServer(us.RPC, us)
	us.StartService()




	/*c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill,syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	s := <-c
	fmt.Println("---service exit----:",s)
	us.StopService()*/
}
