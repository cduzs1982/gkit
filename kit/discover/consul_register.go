package discover

import (
	"fmt"
	api "github.com/hashicorp/consul/api"
	"log"
	"strings"
)

type ConsulRegister struct {
	ip             string
	port           int
	serviceName    string
	monitorAddress string
	client         *api.Client
	id             string
}

func NewConsul(ip string, port int, serviceName string, monitorAddress string) *ConsulRegister {
	return &ConsulRegister{ip: ip, port: port, serviceName: serviceName, monitorAddress: monitorAddress}
}

func (c *ConsulRegister) RegisterService() {

	var tags []string
	c.id = c.serviceName + "-" + c.ip
	service := api.AgentServiceRegistration{
		ID:      c.id,
		Name:    c.serviceName,
		Port:    c.port,
		Address: c.ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			GRPC: fmt.Sprintf("%v:%v/%v", c.ip, c.port, c.serviceName),
			Interval: "5s",
			Timeout:  "1000s",
		},
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	c.client = client
	if err := c.client.Agent().ServiceRegister(&service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul with tags %q", c.serviceName, strings.Join(tags, ","))
}

func (c *ConsulRegister) UnRegisterService() {
	if err := c.client.Agent().ServiceDeregister(c.id); err != nil {
		log.Fatal(err)
	}
}
