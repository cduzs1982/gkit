package discover

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net"
	"strconv"
)


type ConsulDiscover struct {
	Address map[string]string
}

func NewConsulDiscover() *ConsulDiscover {
	return &ConsulDiscover{Address: make(map[string]string)}
}
func (d *ConsulDiscover)DiscoverService(names[] string)  {
	var lastIndex uint64

	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}

	for _, name := range names {
		services, metainfo, err := client.Health().Service(name, "", true, &api.QueryOptions{
			WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
		})
		if err != nil {
			fmt.Printf("error retrieving instances from Consul: %v/n", err)
		}
		lastIndex = metainfo.LastIndex

		for _, service := range services {
			fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
			d.Address[name] = net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
		}
	}
}
