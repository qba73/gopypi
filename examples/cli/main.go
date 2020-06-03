package main

import (
	"fmt"
	"github.com/qba73/gopypi"
)

func main() {
	// Create a new, default PyPI client
	client, err := gopypi.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	// Fetch information about Python "pytest" package
	p, err := client.Package.Get("pytest")
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
