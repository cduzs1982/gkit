package hello_portal

import (
	"awesomeProject/kit"
	pb "awesomeProject/proto/hello"
	"context"
	"log"
)

type HelloPortal struct {
	kit.KitInvoker
	service pb.GreetClient
}

func NewHelloPortal() *HelloPortal  {
	k := kit.NewKitInvoker()

	return &HelloPortal{ KitInvoker:*k, service: pb.NewGreetClient(k.Conn)}
}

func (h *HelloPortal)SayHelloTest(name string)  {
	r, err := h.service.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func (h *HelloPortal)Close()  {
	h.Conn.Close()
}