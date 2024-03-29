package main

import (
	"fmt"

	pypi "github.com/qba73/gopypi"
)

func main() {
	// Use default Get package method to retrieve info about the 'requests' package.
	p, err := pypi.Get("requests")
	if err != nil {
		fmt.Println(err)
	}

	// Print out information about the package.
	fmt.Println(p.Info.Name)
	fmt.Println(p.Info.Version)
	fmt.Println(p.Info.License)
	fmt.Println(p.Info.RequiresPython)

	// Print out classifiers
	for _, c := range p.Info.Classifiers {
		fmt.Println(c)
	}
}
