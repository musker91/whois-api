package configer

import (
	"fmt"
	"path/filepath"
	"whois-api/utils"

	"github.com/jinzhu/configor"
)

var (
	Configer ConfigerStruct
)

func InitialConfier() {
	configerFile := filepath.Join(utils.GetRootPath(), "config.yml")
	err := configor.Load(&Configer, configerFile)
	if err != nil {
		panic(fmt.Sprintf("initial config fail,%s", err))
	}
}
