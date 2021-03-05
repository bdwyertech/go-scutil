package main

import (
	"fmt"
	"log"

	"github.com/bdwyertech/go-scutil/proxy"
)

func main() {
	r, err := proxy.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.ToJSON())
}
