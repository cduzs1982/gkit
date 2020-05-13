package kit

import (
	"net"
)

var lis net.Listener

/*func Init(filename string) *grpc.Server {

	var conf conf
	conf.getConf(filename)

	fmt.Println("server start.....")
	var err error
	lis, err = net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()

	return s
}

func Start(s *grpc.Server) {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve:%v", err)
	}
}*/
