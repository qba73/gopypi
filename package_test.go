package gopypi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func readFile(f string) ([]byte, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func mockServer() *httptest.Server {
	content, err := readFile("testdata/response_pytest_json.json")
	if err != nil {
		log.Fatal(err)
	}

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(content))
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// WithTestClient knows how to use mock server url
// instead of the pypi base url.
func WithTestClient(url string) func(*Client) error {
	return func(c *Client) error {
		c.BaseURL = url
		return nil
	}
}

func TestPackageService_Get(t *testing.T) {
	server := mockServer()
	defer server.Close()

	client, err := NewClient(WithTestClient(server.URL))
	if err != nil {
		t.Fatal("NewClient() error")
	}

	arg := "pytest"
	got, err := client.Package.Get(arg)
	if err != nil {
		t.Errorf("client.Package.Get(%q) = %v", arg, got)
	}

	want := "pytest"
	if got.Info.Name != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", arg, got.Info.Name, want)
	}

	want = "MIT license"
	if got.Info.License != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", arg, got.Info.License, want)
	}

	want = "pytest-5.4.2-py3-none-any.whl"
	if got.URLS[0].Filename != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", arg, got.URLS[0].Filename, want)
	}

	want = "pytest-5.4.2.tar.gz"
	if got.URLS[1].Filename != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", arg, got.URLS[1].Filename, want)
	}
}
