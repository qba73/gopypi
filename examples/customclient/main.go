package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/qba73/gopypi"
)

func main() {
	client := gopypi.NewClient()
	client.HttpClient = &http.Client{}

	p, err := client.Get(context.Background(), "pytest")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(p.Info.Name)
	fmt.Println(p.Info.Version)
	fmt.Println(p.Info.License)
	fmt.Println(p.Info.RequiresPython)

	for _, c := range p.Info.Classifiers {
		fmt.Println(c)
	}
}
