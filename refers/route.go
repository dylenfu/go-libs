package refers

import (
	"github.com/dylenfu/go-libs/refers/ini"
	_ "github.com/dylenfu/go-libs/refers/ini"
)

// 引用时使用_ main入口会运行所有包的init
func Route(sub string) {
	switch sub {
	case "init":
		ini.ExecAllInit()
	}
}
