package main

import (
	"fmt"
	"github.com/Dimadetected/kerniganGoBook/tema4/github"
	"log"
)

func main() {
	search, error := github.GetIssues()
	if error != nil {
		log.Fatal(error.Error())
	}

	fmt.Printf("Всего %d тем: \n", search.TotalCount)
	for _, item := range search.Items {
		fmt.Printf("#%d | %s | %s \n", item.Id, item.Title, item.User.Login)
	}
}
