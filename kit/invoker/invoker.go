package invoker

import (
	"awesomeProject/kit"
	"awesomeProject/kit/discover"
	"google.golang.org/grpc"
	"log"
	"os"
	"path"
)

type KitInvoker struct {
	Conn *grpc.ClientConn
	dis *discover.ConsulDiscover

}

func NewKitInvoker() *KitInvoker {
	var conf kit.ConsumerConf

	dir, _ := os.Getwd()
	conf.GetConsumerConf(path.Join(dir, "kit-consumer.yaml"))

	dis := discover.NewConsulDiscover()
	dis.DiscoverService(conf.Services)

	conn, err := grpc.Dial(dis.Address["user_service"], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &KitInvoker{Conn: conn}
}

func (k *KitInvoker) Close() {
	k.Conn.Close()
}
