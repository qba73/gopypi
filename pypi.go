package gopypi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type PackageUrl struct {
	Filename       string `json:"filename"`
	PackageType    string `json:"package_type"`
	PythonVersion  string `json:"python_version"`
	RequiresPython string `json:"requires_python"`
	UploadTime     string `json:"upload_time"`
	Url            string `json:"url"`
}

type PackageInfo struct {
	Name           string   `json:"name"`
	Classifiers    []string `json:"classifiers"`
	License        string   `json:"license"`
	Version        string   `json:"version"`
	RequiresPython string   `json:"requires_python"`
}

// Package holds all information about Python package.
type Package struct {
	Info PackageInfo  `json:"info"`
	URLS []PackageUrl `json:"urls"`
}

// Client represents PyPI client.
type Client struct {
	BaseURL    string
	HttpClient *http.Client
}

// NewClient creates a new, default PyPI client.
func NewClient() *Client {
	client := Client{
		BaseURL:    "https://pypi.org",
		HttpClient: &http.Client{Timeout: 10 * time.Second},
	}
	return &client
}

// Get knows how to retrieve information from PyPI server.
// It returns info about given Python package.
func (c *Client) Get(ctx context.Context, name string) (Package, error) {
	path := fmt.Sprintf("%s/pypi/%s/json", c.BaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return Package{}, err
	}
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return Package{}, err
	}
	defer res.Body.Close()

	var pkg Package
	if err := json.NewDecoder(res.Body).Decode(&pkg); err != nil {
		return Package{}, err
	}
	return pkg, nil
}

// DefaultPyPIClient it'a a default client used by the default Get function.
var DefaultPyPIClient = NewClient()

// Get takes a string representing a Python package name and returns
// detailed information about the package.
// Get internally uses default Client.
func Get(name string) (Package, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return DefaultPyPIClient.Get(ctx, name)
}

var usage = `Usage: pypi <Python-package-name>`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return 1
	}
	pkgName := os.Args[1]
	pkg, err := Get(pkgName)
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}
	b, err := json.Marshal(pkg)
	if err != nil {
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%+v\n", string(b))
	return 0
}
