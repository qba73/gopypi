package main

import (
	"fmt"
	"github.com/qba73/gopypi/pkg/pypi"
	"log"
)

func main()  {
	client, err := pypi.NewClient()
	if err != nil {
		log.Println(err)
	}

	p, err := client.Package.Get("pytest")
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
