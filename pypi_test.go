package gopypi_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/qba73/gopypi"
)

func readFile(filepath string) ([]byte, error) {
	b, err := os.ReadFile(filepath)
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

func TestPackageService_Get(t *testing.T) {
	server := mockServer()
	defer server.Close()

	client := gopypi.NewClient()
	client.BaseURL = server.URL

	arg := "pytest"
	got, err := client.Get(context.Background(), arg)
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
