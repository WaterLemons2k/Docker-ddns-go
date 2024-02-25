package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/WaterLemons2k/Docker-ddns-go/tools/json"
	"github.com/WaterLemons2k/Docker-ddns-go/tools/untar"
)

// Release represents a GitHub release.
//
// https://docs.github.com/en/rest/releases/releases#get-a-release-by-tag-name
type Release struct {
	Assets []Asset `json:"assets"`
}

// Asset represents a GitHub release asset.
type Asset struct {
	Name string `json:"name"`
	URL  string `json:"browser_download_url"`
}

const (
	// assetName indicates that the file with assetName in the asset will be downloaded.
	assetName = "linux_x86_64"

	// repo is the name of the GitHub repository without the owner part.
	repo = "ddns-go"
)

var client = &http.Client{Timeout: 10 * time.Second}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run -C tools . <version>")
		return
	}
	version := os.Args[1]

	url := fmt.Sprintf("https://api.github.com/repos/jeessy2/%s/releases/tags/v%s", repo, version)

	body := getBody(url)
	defer body.Close()

	var release Release
	if err := json.Parse(body, &release); err != nil {
		log.Fatal(err)
	}

	for _, asset := range release.Assets {
		// We only focus on asset with assetName
		if !strings.Contains(asset.Name, assetName) {
			continue
		}

		body := getBody(asset.URL)
		defer body.Close()
		fmt.Println(asset.Name, "downloaded. Extracting", repo, "from the tar.gz file...")

		// Extracts the file with repo from the tar.gz file
		untar.SpecificFile(body, repo)
		break
	}
}

// getBody gets the response body from the given url.
func getBody(url string) io.ReadCloser {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}
