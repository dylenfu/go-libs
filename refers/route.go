package refers

import (
	_ "github.com/dylenfu/go-libs/refers/ini"
	"github.com/dylenfu/go-libs/refers/ini"
)

func Route(sub string) {
	switch sub {
	case "init":
		ini.ExecAllInit()
	}
}
