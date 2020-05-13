package kit

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type ProviderConf struct {
	Register struct {
		IP string `yaml:"ip"`
		Port int `yaml:"port"`
		ServiceName    string `yaml:"serviceName"`
		MonitorAddress string `yaml:"monitorAddress"`
	}
	Port string `yaml:"port"`
}

func (c *ProviderConf) GetConf(filename string) *ProviderConf {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

type ConsumerConf struct {
	Address string `yaml:"address"`
	Services []string
}

func (c *ConsumerConf)GetConsumerConf(filename string) *ConsumerConf  {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
