package main

import (
	"fmt"

	"github.com/farrukhny/auth/common"
)

func main() {
	common.StartUp()
	fmt.Println(common.AppConfig.DBHost)
}
