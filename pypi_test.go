package gopypi_test

import (
	"context"
	"fmt"
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

func testServer(t *testing.T) *httptest.Server {
	t.Helper()

	content, err := readFile("testdata/response_pytest_json.json")
	if err != nil {
		t.Fatal(err)
	}

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(content))
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestPackageService_Get(t *testing.T) {
	t.Parallel()

	server := testServer(t)
	defer server.Close()

	client := gopypi.NewClient()
	client.BaseURL = server.URL

	pythonPackage := "pytest"
	got, err := client.Get(context.Background(), pythonPackage)
	if err != nil {
		t.Errorf("client.Package.Get(%q) = %v", pythonPackage, got)
	}

	want := "pytest"
	if got.Info.Name != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", pythonPackage, got.Info.Name, want)
	}

	want = "MIT license"
	if got.Info.License != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", pythonPackage, got.Info.License, want)
	}

	want = "pytest-5.4.2-py3-none-any.whl"
	if got.URLS[0].Filename != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", pythonPackage, got.URLS[0].Filename, want)
	}

	want = "pytest-5.4.2.tar.gz"
	if got.URLS[1].Filename != want {
		t.Errorf("client.Package.Get(%q) = %s; want %s", pythonPackage, got.URLS[1].Filename, want)
	}
}
