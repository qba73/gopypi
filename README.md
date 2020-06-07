# gopypi

[![CircleCI](https://circleci.com/gh/qba73/gopypi.svg?style=shield)](https://circleci.com/gh/qba73/gopypi)
[![codecov](https://codecov.io/gh/qba73/gopypi/branch/master/graph/badge.svg)](https://codecov.io/gh/qba73/gopypi)


Go client for [PyPI](https://pypi.org) REST API.


# Installation
```go
$ go get github.com/qba73/gopypi
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/qba73/gopypi"
)

func main() {
	// Create a new, default PyPI client
	pypi, err := gopypi.NewClient()
	if err != nil {
		fmt.Println(err)
	}

	// Fetch information about Python "pytest" package
	p, err := pypi.Package.Get("pytest")
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
```


