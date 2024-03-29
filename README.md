[![CircleCI](https://circleci.com/gh/qba73/gopypi.svg?style=shield)](https://circleci.com/gh/qba73/gopypi)
[![codecov](https://codecov.io/gh/qba73/gopypi/branch/master/graph/badge.svg)](https://codecov.io/gh/qba73/gopypi)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/gopypi)](https://goreportcard.com/report/github.com/qba73/gopypi)

`gopypi` is a Go client library for [PyPI](https://pypi.org) REST API.

## How to use it

### As a package in your project

Import package

```go
import "github.com/qba73/gopypi"

func main() {
 package, err := gopypi.Get("requests")
 if err != nil {
  // handle error
 }
 fmt.Println(package)
}
```

### As a cli tool

Build the binary:

```shell
go build -o pypi ./cmd/pypi/main.go
```

Run cli:

```shell
./pypi requests
```

```shell
{"info":{"name":"requests","classifiers":["Development Status :: 5 - Production/Stable","Environment :: Web Environment","Intended Audience :: Developers","License :: OSI Approved :: Apache Software License","Natural Language :: English","Operating System :: OS Independent","Programming Language :: Python","Programming Language :: Python :: 3","Programming Language :: Python :: 3 :: Only","Programming Language :: Python :: 3.10","Programming Language :: Python :: 3.11","Programming Language :: Python :: 3.7","Programming Language :: Python :: 3.8","Programming Language :: Python :: 3.9","Programming Language :: Python :: Implementation :: CPython","Programming Language :: Python :: Implementation :: PyPy","Topic :: Internet :: WWW/HTTP","Topic :: Software Development :: Libraries"],"license":"Apache 2.0","version":"2.31.0","requires_python":"\u003e=3.7"},"urls":[{"filename":"requests-2.31.0-py3-none-any.whl","package_type":"","python_version":"py3","requires_python":"\u003e=3.7","upload_time":"2023-05-22T15:12:42","url":"https://files.pythonhosted.org/packages/70/8e/0e2d847013cb52cd35b38c009bb167a1a26b2ce6cd6965bf26b47bc0bf44/requests-2.31.0-py3-none-any.whl"},{"filename":"requests-2.31.0.tar.gz","package_type":"","python_version":"source","requires_python":"\u003e=3.7","upload_time":"2023-05-22T15:12:44","url":"https://files.pythonhosted.org/packages/9d/be/10918a2eac4ae9f02f6cfe6414b7a155ccd8f7f9d4380d62fd5b955065c3/requests-2.31.0.tar.gz"}]}
```

Pipe output to `jq`:

```shell
./pypi requests | jq .
{
  "info": {
    "name": "requests",
    "classifiers": [
      "Development Status :: 5 - Production/Stable",
      "Environment :: Web Environment",
      "Intended Audience :: Developers",
      "License :: OSI Approved :: Apache Software License",
      "Natural Language :: English",
      "Operating System :: OS Independent",
      "Programming Language :: Python",
      "Programming Language :: Python :: 3",
      "Programming Language :: Python :: 3 :: Only",
      "Programming Language :: Python :: 3.10",
      "Programming Language :: Python :: 3.11",
      "Programming Language :: Python :: 3.7",
      "Programming Language :: Python :: 3.8",
      "Programming Language :: Python :: 3.9",
      "Programming Language :: Python :: Implementation :: CPython",
      "Programming Language :: Python :: Implementation :: PyPy",
      "Topic :: Internet :: WWW/HTTP",
      "Topic :: Software Development :: Libraries"
    ],
    "license": "Apache 2.0",
    "version": "2.31.0",
    "requires_python": ">=3.7"
  },
  "urls": [
    {
      "filename": "requests-2.31.0-py3-none-any.whl",
      "package_type": "",
      "python_version": "py3",
      "requires_python": ">=3.7",
      "upload_time": "2023-05-22T15:12:42",
      "url": "https://files.pythonhosted.org/packages/70/8e/0e2d847013cb52cd35b38c009bb167a1a26b2ce6cd6965bf26b47bc0bf44/requests-2.31.0-py3-none-any.whl"
    },
    {
      "filename": "requests-2.31.0.tar.gz",
      "package_type": "",
      "python_version": "source",
      "requires_python": ">=3.7",
      "upload_time": "2023-05-22T15:12:44",
      "url": "https://files.pythonhosted.org/packages/9d/be/10918a2eac4ae9f02f6cfe6414b7a155ccd8f7f9d4380d62fd5b955065c3/requests-2.31.0.tar.gz"
    }
  ]
}
```
