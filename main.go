package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/jabbalaci/szakdoga/lib/jweb"
)

func fetchHTML(url string) string {
	fmt.Println("# downloading the webpage...")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("fetchHTML: bad status: %s", resp.Status))
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func extractPDFURL(html string) string {
	re := regexp.MustCompile(`<meta name="citation_pdf_url" content="([^"]+)">`)
	match := re.FindStringSubmatch(html)
	if len(match) != 2 {
		panic("extractPDFURL: no match")
	}

	return match[1]
}

func readURL() string {
	fmt.Println("A hallgatói dolgozatok (Informatikai Kar) itt érhetők el: https://dea.lib.unideb.hu")
	fmt.Println("Egy szakdolgozat URL-je így néz ki (példa): https://dea.lib.unideb.hu/items/ed260496-92b4-428e-a8a5-5c9bd0c0f28f")
	fmt.Println()

	var urlStr string
	fmt.Print("A letöltendő szakdolgozat URL-je: ")
	fmt.Scanln(&urlStr)

	return urlStr
}

func main() {
	urlStr := readURL()
	html := fetchHTML(urlStr)
	pdfURL := extractPDFURL(html)
	fmt.Println("# URL of the PDF:")
	fmt.Println("#", pdfURL)
	fmt.Println("# opening the PDF in your web browser...")
	err := jweb.OpenInBrowser(pdfURL)
	if err != nil {
		panic(err)
	}
}
