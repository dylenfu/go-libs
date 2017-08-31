package toml

import (
	"net"
	"time"
	"os"
	"github.com/naoina/toml"
	"log"
)

type tomlConfig struct {
	Title string
	Owner struct {
		Name string
		Dob time.Time
	}

	Database struct {
		Server string
		Ports []int
		ConnectionMax int
		Enabled bool
	}

	Servers map[string]ServerInfo

	Clients struct {
		Data [][]interface{}
		Hosts []string
	}
}

type ServerInfo struct {
	Ip net.IP
	Dc string
}

func QuickStart() {
	dir, _ := os.Getwd()
	f, err := os.Open(dir + "/toml/qs.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var conf tomlConfig
	if err := toml.NewDecoder(f).Decode(&conf); err != nil {
		panic(err)
	}
	log.Println(conf)
	log.Println(conf.Servers["alpha"].Ip)
	log.Println(conf.Clients.Data[0][0])
}