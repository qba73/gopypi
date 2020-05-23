package pypi

import (
	"encoding/json"
	"fmt"
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

type Package struct {
	Info PackageInfo  `json:"info"`
	URLS []PackageUrl `json:"urls"`
}

type PackageService struct {
	client *Client
}

func (p *PackageService) Get(name string) (*Package, error) {
	path := fmt.Sprintf("/pypi/%s/json", name)
	res, err := p.client.httpClient.Get(p.client.BaseURL + path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pkg Package
	if err := json.NewDecoder(res.Body).Decode(&pkg); err != nil {
		return nil, err
	}
	return &pkg, nil
}
