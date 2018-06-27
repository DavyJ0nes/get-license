package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/davyj0nes/get-license/htmlparser"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

func main() {
	license := flag.String("license", "", "license to get")
	flag.Usage = func() {
		fmt.Printf("Usage: getlicense [options]\n\n")
		flag.PrintDefaults()
		fmt.Println(`List of Available Licenses:
	- apache
	- gpl
	- mit
	- mozilla
	`)
	}

	flag.Parse()

	URL, err := getLicenseURL(*license)
	if err != nil {
		panic(err)
	}

	licenseContent, err := getLicense(URL)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, licenseContent)
}

// getLicense gets a string representation of the required license file
func getLicense(URL string) (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", errors.Wrap(err, "Error Calling URL")
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Error Parsing HTML")
	}

	licenseNode := htmlparser.GetElementByID(doc, "license-text")
	licenseText := licenseNode.FirstChild.Data

	return licenseText, nil
}

// getLicenseURL is a simple parser taking user arguments and returning a URL.
func getLicenseURL(license string) (string, error) {
	baseURL := "https://choosealicense.com/licenses/"

	license = strings.ToLower(license)

	switch license {
	case "apache":
		return baseURL + "apache-2.0", nil
	case "mit":
		return baseURL + "mit", nil
	case "mozilla":
		return baseURL + "mpl-2.0", nil
	case "gpl":
		return baseURL + "gpl-3.0", nil
	case "":
		return "", fmt.Errorf("no license provided")
	default:
		return "", fmt.Errorf("unknown license: %s", license)
	}
}
