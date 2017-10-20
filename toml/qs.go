package toml

import (
	"net"
	"time"
	"os"
	"github.com/naoina/toml"
	"log"
	"math/big"
)

type ServerInfo struct {
	Ip net.IP
	Dc string
}

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

////////////////////////////////////////////////////////////////////////
// unmarshal
////////////////////////////////////////////////////////////////////////

type RawTOML []byte

func (r *RawTOML) UnmarshalTOML(input []byte) error {
	cpy := make([]byte, len(input))
	copy(cpy, input)
	*r = cpy
	return nil
}

func SimpleUnmarshal() {
	println("hi")
	input := []byte(`
		foo = 10000000000000000000000000000000000000000000000000999999999999999999999999999993333333333333

		[[servers]]
		addr = "198.51.100.3:80" # a comment

		[[servers]]
		addr = "192.0.2.10:8080"
		timeout = "30s"
		`)

	var config struct {
		Foo     *big.Int
		Servers RawTOML
	}

	toml.Unmarshal(input, &config)
	log.Println("config.Foo = ", config.Foo.String())
	log.Println("config.servers = ", string(config.Servers))
	//fmt.Printf("config.Servers =\n%s\n", indent(config.Servers, 2))
}
