package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farrukhny/auth/common"
)

func init() {
	common.StartUp()
}

func main() {

	log.Print(common.AppConfig)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(common.AppConfig.Port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	//log.Println(&common.AppConfig.AppName)

	//dd := &common.AppConfig.AppName

	log.Printf("%+v", common.AppConfig.AppName)
	//log.Println(*dd)
	fmt.Fprintf(w, "%v", common.AppConfig.AppName)
}
