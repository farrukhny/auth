package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/farrukhny/auth/models"

	"github.com/farrukhny/auth/common"
)

func init() {
	common.StartUp()
}

func main() {
	http.HandleFunc("/", Do)
	http.HandleFunc("/again", DoAgain)
	log.Fatal(http.ListenAndServe(common.AppConfig.Port, nil))
}

func DoAgain(w http.ResponseWriter, r *http.Request) {
	book := models.IndexDB()
	for _, bk := range book {
		fmt.Fprintf(w, "%s, %s, %s, %.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}

func Do(w http.ResponseWriter, r *http.Request) {
	dd := "Hello"
	fmt.Fprintln(w, dd)
}
