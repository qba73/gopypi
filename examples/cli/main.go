package main

import (
	"context"
	"fmt"

	pypi "github.com/qba73/gopypi"
)

func main() {
	// Create a new, default PyPI client
	client := pypi.NewClient()

	// Fetch information about Python "pytest" package
	p, err := client.Get(context.Background(), "pytest")
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
