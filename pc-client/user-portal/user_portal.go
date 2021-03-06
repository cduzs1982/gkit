package user_portal

import (
	"awesomeProject/kit/invoker"
	pb "awesomeProject/proto/user"
	"context"
	"log"
)


type UserPortal struct {
	invoker.KitInvoker

	service pb.UserClient
}

func NewUserPortal() *UserPortal {
	k := invoker.NewKitInvoker()

	return &UserPortal{ KitInvoker:*k, service: pb.NewUserClient(k.Conn)}
}

func (u *UserPortal)UserLogin()  {
	rsp, err := u.service.Login(context.Background(), &pb.LoginRequest{Name: " pc-client"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("user login: %s", rsp.Message)

}

func (u *UserPortal)Close() {
	u.Conn.Close()

}