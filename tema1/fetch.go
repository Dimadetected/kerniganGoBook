package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	params := os.Args[1:]
	for _, param := range params {
		if strings.HasPrefix(param, "https://") == false {
			param = "https://" + param
		}
		request, err := http.Get(param)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(param, request.Status)
	}
}
